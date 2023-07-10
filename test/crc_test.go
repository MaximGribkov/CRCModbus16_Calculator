package test

import (
	"com_with_go/pkg"
	"testing"
)

func TestCalculateCRC(t *testing.T) {
	tests := []struct {
		data     []byte
		expected []byte
	}{
		{[]byte{0x00, 0x01, 0x02, 0x03}, []byte{0x10, 0x85}},
		{[]byte{0xFF, 0xFF, 0xFF, 0xFF}, []byte{0x01, 0xB0}},
		{[]byte{0x00, 0x00, 0x00, 0x00}, []byte{0x00, 0x24}},
		{[]byte{0x02, 0x03, 0x01, 0xF4, 0x00, 0x02}, []byte{0x84, 0x36}},
	}

	for _, test := range tests {
		result := pkg.CalculateCRC(test.data)
		if len(result) != len(test.expected) {
			t.Errorf("Incorrect result length for data %v. Expected %v, got %v", test.data, test.expected, result)
		}
		for i := 0; i < len(result); i++ {
			if result[i] != test.expected[i] {
				t.Errorf("Incorrect result for data %v. Expected %v, got %v", test.data, test.expected, result)
				break
			}
		}
	}
}
