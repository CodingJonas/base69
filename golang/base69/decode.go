package base69

import (
	"strings"
)

func Decode(value []rune) string {
	// return decodeChunk(value, numPadding)
	numPadding := 0
	var sb strings.Builder
	for i := 0; i+16 <= len(value); i += 16 {
		runeChunk := value[i : i+16]
		encodedChunk := ""
		// Check if we decode the last chunk, in this case remove the padding
		if i == len(value)-16 {
			numPadding = getNumPadding(value)
		}
		encodedChunk = decodeChunk(runeChunk, numPadding)
		sb.WriteString(encodedChunk)
	}
	return sb.String()
}

func getNumPadding(value []rune) int {
	if value[len(value)-1] == '=' {
		return int(value[len(value)-2] - '0')
	}
	return 0
}

// Each chunk must contain exactly 16 runes
func decodeChunk(runeChunk []rune, numPadding int) string {
	nums := getNumsFromChunk(runeChunk)
	bytes := transformNumsToBytes(nums)
	return string(bytes[:len(bytes)-numPadding])
}

// Transform the original input runes into numerical values later on used to create the original bits
func getNumsFromChunk(runeChunk []rune) [8]uint8 {
	nums := [8]uint8{}
	for i := 0; i <= 7; i += 1 {
		s := runeChunk[i*2 : i*2+2]
		nums[i] = getNumFromBase69(s)
	}
	return nums
}

// Transform numbers back into original byte array
func transformNumsToBytes(nums [8]uint8) [7]byte {
	bytes := [7]byte{}
	// Just a table to help make sense out of the bit operations
	// 0        1       2       3       4       5       6       7
	// .XXXXXXX.XXXXXXX.XXXXXXX.XXXXXXX.XXXXXXX.XXXXXXX.XXXXXXX.XXXXXXX
	//  0000000.0111111.1122222.2223333.3333444.4444455.5555556.6666666
	bytes[0] = nums[0]<<1 + nums[1]&0b01000000>>6
	bytes[1] = nums[1]<<2 + nums[2]&0b01100000>>5
	bytes[2] = nums[2]<<3 + nums[3]&0b01110000>>4
	bytes[3] = nums[3]<<4 + nums[4]&0b01111000>>3
	bytes[4] = nums[4]<<5 + nums[5]&0b01111100>>2
	bytes[5] = nums[5]<<6 + nums[6]&0b01111110>>1
	bytes[6] = nums[6]<<7 + nums[7]&0b01111111
	return bytes
}

func getNumFromBase69(s []rune) uint8 {
	remainderRune, quotientRune := s[0], s[1]
	quotient := strings.Index(CHARS, string(quotientRune))
	remainder := strings.Index(CHARS, string(remainderRune))
	return uint8(quotient*69 + remainder)
}
