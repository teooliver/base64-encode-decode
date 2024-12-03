package main

import "testing"

func TestPadChunk(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "000000"},
		{"1", "100000"},
		{"11", "110000"},
		{"111", "111000"},
		{"1111", "111100"},
		{"11111", "111110"},
		{"111111", "111111"},
	}

	for _, test := range tests {
		result := padChunk(test.input)
		if result != test.expected {
			t.Errorf("padChunk(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestConvert6bitTo8bit(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"000000", "00000000"},
		{"000001", "00000001"},
		{"111111", "00111111"},
		{"101010", "00101010"},
		{"010101", "00010101"},
	}

	for _, test := range tests {
		result := convert6bitTo8bit(test.input)
		if result != test.expected {
			t.Errorf("convert6bitTo8bit(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestStringToBin(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"A", "1000001"},
		{"B", "1000010"},
		{"C", "1000011"},
		{"a", "1100001"},
		{"b", "1100010"},
		{"c", "1100011"},
		{"1", "110001"},
		{"2", "110010"},
		{"3", "110011"},
	}

	for _, test := range tests {
		result := stringToBin(test.input)
		if result != test.expected {
			t.Errorf("stringToBin(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}
