package base69

import "testing"

func TestDecode_abcdefg(t *testing.T) {
	input := []rune("wATBHB2AjAVAHBiB")
	expected := "abcdefg"

	actual := Decode(input)
	if actual != expected {
		t.Fatalf(`Got "%s", but expected "%s"`, actual, expected)
	}
}

func TestDecode_abcdefghijklmn(t *testing.T) {
	input := []rune("wATBHB2AjAVAHBiB0AaAtAmAWBxAVBpB")
	expected := "abcdefghijklmn"

	actual := Decode(input)
	if actual != expected {
		t.Fatalf(`Got "%s", but expected "%s"`, actual, expected)
	}
}

func TestDecode_sixtynine(t *testing.T) {
	input := []rune("5AVBvAHAjAgBXBkB3AZAkAQAAAAAAA4=")
	expected := "sixtynine!"

	actual := Decode(input)
	if actual != expected {
		t.Fatalf(`Got "%s", but expected "%s"`, actual, expected)
	}
}

func TestDecode_abcd(t *testing.T) {
	input := []rune("wATBHB2AgAAAAA3=")
	expected := "abcd"

	actual := Decode(input)
	if actual != expected {
		t.Fatalf(`Got "%s", but expected "%s"`, actual, expected)
	}
}

func TestDecode_long(t *testing.T) {
	input := []rune("qAaANAXAZABANBuBQAYAkAGAeB9AXBiBQAdAMASBbALBXBgArAVBuABB*AAAPBqA5AXBMARBbAlA<AnBVAKA|AHAbAgBVBdB3AWBOAyABAjBNB1=")
	expected := "This is a long test. With **special** symbols :)"

	actual := Decode(input)
	if actual != expected {
		t.Fatalf(`Got "%s", but expected "%s"`, actual, expected)
	}
}

func TestGetNumsFromChunk(t *testing.T) {
	input := []rune("wATBHB2AjAVAHBiB")
	expected := [8]uint8{48, 88, 76, 54, 35, 21, 76, 103}

	actual := getNumsFromChunk(input)
	if !isUint8SlicesEqual(actual, expected) {
		t.Fatalf(`Got "%v", but expected "%v"`, actual, expected)
	}
}

func TestTransformNumsToBytes(t *testing.T) {
	input := [8]uint8{48, 88, 76, 54, 35, 21, 76, 103}
	expected := [7]byte{97, 98, 99, 100, 101, 102, 103}

	actual := transformNumsToBytes(input)
	if !isByteSlicesEqual(actual[:], expected[:]) {
		t.Fatalf(`Got "%v", but expected "%v"`, actual, expected)
	}
}

func TestGetNumPadding(t *testing.T) {
	input := []rune("wATBHB2AgAAAAA3=")
	expected := 3

	actual := getNumPadding(input)
	if actual != expected {
		t.Fatalf(`Got "%v", but expected "%v"`, actual, expected)
	}
}

func isUint8SlicesEqual(a, b [8]uint8) bool {
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
