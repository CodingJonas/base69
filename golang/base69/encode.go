package base69

import (
	"strings"
)

const CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/-*<>|"

func Encode(bytes []byte) string {
	// Ensure string is a multiple of 7 bytes long
	bytes, numPaddingBytes := padBytes(bytes)
	var sb strings.Builder
	for i := 0; i+7 <= len(bytes); i += 7 {
		byteChunk := bytes[i : i+7]
		encodedChunk := encodeChunk(byteChunk)
		sb.WriteString(encodedChunk)
	}
	return appendOutputPadding(sb.String(), numPaddingBytes)
}

func padBytes(bytes []byte) ([]byte, int) {
	extraBytes := len(bytes) % 7
	if extraBytes != 0 {
		numPaddingBytes := 7 - extraBytes
		paddingBytes := getPaddingSlice(numPaddingBytes)
		bytes := append(bytes, paddingBytes...)
		return bytes, numPaddingBytes
	}
	return bytes, 0
}

func getPaddingSlice(numPaddingBytes int) []byte {
	paddingBytes := make([]byte, numPaddingBytes)
	for i := range numPaddingBytes {
		paddingBytes[i] = 0b00000000
	}
	return paddingBytes
}

func appendOutputPadding(s string, numPaddingBytes int) string {
	if numPaddingBytes == 0 {
		// No padding took place
		return s
	}
	// We can safely overwrite the last two characters to indicate padding, they are from the padded 0s only
	r := []rune(s)
	r[len(r)-2] = rune('0' + numPaddingBytes)
	r[len(r)-1] = '='
	return string(r)
}

// Each chunk must contain exactly 7 byte.
func encodeChunk(byteChunk []byte) string {
	nums := [8]uint8{}
	// Just a table to help make sense out of the bit operations
	// 0        1        2        3        4        5        6
	// XXXXXXXX.XXXXXXXX.XXXXXXXX.XXXXXXXX.XXXXXXXX.XXXXXXXX.XXXXXXXX
	// 00000001.11111122.22222333.33334444.44455555.55666666.67777777
	nums[0] = byteChunk[0] >> 1
	nums[1] = byteChunk[0]&0b1<<6 + byteChunk[1]>>2
	nums[2] = byteChunk[1]&0b11<<5 + byteChunk[2]>>3
	nums[3] = byteChunk[2]&0b111<<4 + byteChunk[3]>>4
	nums[4] = byteChunk[3]&0b1111<<3 + byteChunk[4]>>5
	nums[5] = byteChunk[4]&0b11111<<2 + byteChunk[5]>>6
	nums[6] = byteChunk[5]&0b111111<<1 + byteChunk[6]>>7
	nums[7] = byteChunk[6] & 0b1111111

	var sb strings.Builder
	for _, num := range nums {
		sb.WriteString(getBase69FromNum(num))
	}
	return sb.String()
}

func getBase69FromNum(num uint8) string {
	remainder, quotient := num%69, num/69
	return string([]byte{CHARS[remainder], CHARS[quotient]})
}
