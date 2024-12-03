package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	result := b64Encode("hello")

	fmt.Println(result)

}

func b64Encode(s string) string {
	chars := strings.Split(s, "")
	// Convert each character to binary
	binArray := []string{}
	for _, char := range chars {
		binArray = append(binArray, stringToBin(char))
	}

	// Concatenate the binary strings
	concatChars := strings.Join(binArray, "")

	// Divide the concatenated binary string into 6 bit chunks
	chunks := []string{}
	for i := 0; i < len(concatChars); i += 6 {
		end := i + 6
		if end > len(concatChars) {
			end = len(concatChars)
		}
		chunks = append(chunks, concatChars[i:end])
	}

	// Pad the last chunk with zeros if len less then 6
	chunks[len(chunks)-1] = padChunk(chunks[len(chunks)-1])

	// Convert each 6 bit chunk to an 8 bit chunk by adding "00" to the front
	for i, chunk := range chunks {
		chunks[i] = convert6bitTo8bit(chunk)
	}

	base64String := []string{}
	for _, chunk := range chunks {
		// Convert the 8 bit chunk to decimal. Find the corresponding decimal value in the ASCII table
		decimal := convert8bitToDecimal(chunk)
		base64String = append(base64String, getLetterFromBase64Indice(decimal))
	}

	// Pad the base64 string with "=" if the length is not a multiple of 4
	return padBase64(strings.Join(base64String, ""))

}

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

// Pad last chunk with zeros
func padChunk(chunk string) string {
	for len(chunk) < 6 {
		chunk = fmt.Sprintf("%s%s", chunk, "0")
	}
	return chunk
}

func convert6bitTo8bit(chunk string) string {
	return fmt.Sprintf("00%s", chunk)
}

type asciiTable struct {
	Decimal int    `json:"decimal"`
	Binary  string `json:"binary"`
	// hex     string
	// name    string
}

func convert8bitToDecimal(chunk string) int {

	var asciiTable []asciiTable
	file, err := os.Open("ascii_table.json")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	json.Unmarshal(byteValue, &asciiTable)

	for _, ascii := range asciiTable {
		if ascii.Binary == chunk {
			return ascii.Decimal
		}
	}

	// Return -1 if no match is found because 0 is a valid decimal value and has meaning in this context
	return -1
}

func padBase64(b64 string) string {
	for len(b64)%4 != 0 {
		b64 += "="
	}
	return b64
}

// TODO: Should instead read from file?
func getLetterFromBase64Indice(base64Indice int) string {
	base64Map := map[int]string{
		0:  "A",
		1:  "B",
		2:  "C",
		3:  "D",
		4:  "E",
		5:  "F",
		6:  "G",
		7:  "H",
		8:  "I",
		9:  "J",
		10: "K",
		11: "L",
		12: "M",
		13: "N",
		14: "O",
		15: "P",
		16: "Q",
		17: "R",
		18: "S",
		19: "T",
		20: "U",
		21: "V",
		22: "W",
		23: "X",
		24: "Y",
		25: "Z",
		26: "a",
		27: "b",
		28: "c",
		29: "d",
		30: "e",
		31: "f",
		32: "g",
		33: "h",
		34: "i",
		35: "j",
		36: "k",
		37: "l",
		38: "m",
		39: "n",
		40: "o",
		41: "p",
		42: "q",
		43: "r",
		44: "s",
		45: "t",
		46: "u",
		47: "v",
		48: "w",
		49: "x",
		50: "y",
		51: "z",
		52: "0",
		53: "1",
		54: "2",
		55: "3",
		56: "4",
		57: "5",
		58: "6",
		59: "7",
		60: "8",
		61: "9",
		62: "+",
		63: "/",
	}

	return base64Map[base64Indice]
}
