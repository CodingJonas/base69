package base69

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	input := []byte("")
	expected := ""

	actual := Encode(input)
	if actual != expected {
		t.Fatalf(`Got "%s", but expected "%s"`, actual, expected)
	}
}

func Test_abcd(t *testing.T) {
	input := []byte("abcd")
	expected := "wATBHB2AgAAAAA3="

	actual := Encode(input)
	if actual != expected {
		t.Fatalf(`Got "%s", but expected "%s"`, actual, expected)
	}
}

func Test_abcdefgabcdefghijklmn(t *testing.T) {
	input := []byte("abcdefghijklmn")
	expected := "wATBHB2AjAVAHBiB0AaAtAmAWBxAVBpB"

	actual := Encode(input)
	if actual != expected {
		t.Fatalf(`Got "%s", but expected "%s"`, actual, expected)
	}
}

func Test_long(t *testing.T) {
	input := []byte("This is a long test. With **special** symbols :)")
	expected := "qAaANAXAZABANBuBQAYAkAGAeB9AXBiBQAdAMASBbALBXBgArAVBuABB*AAAPBqA5AXBMARBbAlA<AnBVAKA|AHAbAgBVBdB3AWBOAyABAjBNB1="

	actual := Encode(input)
	if actual != expected {
		t.Fatalf(`Got "%s", but expected "%s"`, actual, expected)
	}
}

func TestEncodeChunk(t *testing.T) {
	// These 7 characters produce a byte array of exactly 7 byte
	input := []byte("abcdefg")
	expected := "wATBHB2AjAVAHBiB"

	actual := encodeChunk(input)
	if actual != expected {
		t.Fatalf(`Got "%s", but expected "%s"`, actual, expected)
	}
}

func TestPadBytes(t *testing.T) {
	input := []byte{0b10101010, 0b11001100, 0b11110001, 0b00001111, 0b11111111}
	expected := []byte{0b10101010, 0b11001100, 0b11110001, 0b00001111, 0b11111111, 0b00000000, 0b00000000}

	actual, _ := padBytes(input)
	if !isByteSlicesEqual(actual, expected) {
		t.Fatal("Byte slice was not properly padded to a multiple of 7")
	}

}

func isByteSlicesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
