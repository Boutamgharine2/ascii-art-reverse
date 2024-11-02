package asci

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	cannotBe map[string]bool
	good     = [][8]string{}
)

// Reverse processes the input ASCII art from a file if it starts with the prefix "--reverse="
func Reverse(s string) bool {
	cannotBe = make(map[string]bool)
	cannotBe["shadow"] = true
	cannotBe["thinkertoy"] = true
	cannotBe["standard"] = true

	if !strings.HasPrefix(s, "--reverse=") {
		return false
	}
	s = strings.TrimPrefix(s, "--reverse=")
	file, err := os.ReadFile(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := strings.Split(string(file), "\n")
	if len(data)-1 == len(file) {
		for i := 1; i < len(data); i++ {
			fmt.Print("\\n")
		}
		return true
	}
	counter := 0
	var g []bool

	for i, line := range data {
		if counter == 8 {
			counter = 0
		}
		if counter == 0 && i == len(data)-1 {
			if len(g) == 0 {
				break
			}
			good = append(good, addCharacter(g, [8]string(data[i-8:i]))...)
			break
		}

		if line == "" && i < len(data)-1 {
			if counter != 0 {
				fmt.Printf("ERROR: Malformed file.\nLine %v should not be an empty line.\n", i)
				os.Exit(0)
			}
			if len(g) > 0 {
				good = append(good, addCharacter(g, [8]string(data[i-8:i]))...)
			}
			good = append(good, [8]string{"\n", "", "", "", "", "", "", ""})

			g = make([]bool, len(line))
			continue
		}

		counter++
		if counter == 1 {
			if i > 7 {
				good = append(good, addCharacter(g, [8]string(data[i-8:i]))...)
			}
			g = make([]bool, len(line))
		}

		if len(g) != len(line) {
			fmt.Printf("ERROR: Malformed file.\nLine %v should be %v characters long.\n", i, len(g))
			os.Exit(0)
		}

		for i1, c := range line {
			if c != ' ' {
				g[i1] = true
				handleCharacter(c)
			}
		}
	}

	reverseASCII()
	return true
}

// addCharacter combines character data into structured format
func addCharacter(g []bool, data [8]string) [][8]string {
	if len(data[0]) == 0 {
		return [][8]string{{"\n", "", "", "", "", "", "", ""}}
	}
	character := [8]string{}
	var lines [][8]string

	for i1, c := range g {
		if !c {
			if character[0] != "" {
				lines = append(lines, character)
			} else {
				lines = append(lines, [8]string{" ", "", "", "", "", "", "", ""})
			}
			character = [8]string{}
		} else {
			for j := 0; j < 8; j++ {
				ln := data[j]
				s := string(ln[i1])
				character[j] += s
			}
		}
	}
	lines = append(lines, [8]string{"\n", "", "", "", "", "", "", ""})
	return lines
}

// reverseASCII processes the ASCII art to generate the final output
func reverseASCII() {
	if len(good) == 0 {
		fmt.Println("good = 0")
		return
	}
	good = good[:len(good)-1]
	results := []string{}
	for n, v := range cannotBe {
		var file []byte
		res := ""
		if v {
			// fmt.Print(n, v)
			file1, err := os.ReadFile(n + ".txt")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			file = file1
			if n == "thinkertoy" {
				file = bytes.ReplaceAll(file1, []byte("\r\n"), []byte("\n"))
			}
		} else {
			continue
		}
		counter := 0
		for _, v := range good {
			rgx := "\n\n"
			if strings.Join(v[:], "") == " " {
				counter++
				if counter == 6 {
					res += " "
					counter = 0
				}
				continue
			} else if strings.Join(v[:], "") == "\n" {
				res += "\\n"
				continue
			}
			endline := "\n"
			for _, v1 := range v {
				rgx += regexp.QuoteMeta(v1) + "\\s+" + endline
			}

			if len(rgx) > 2 {
				regex := regexp.MustCompile(rgx)
				indices := regex.FindAllIndex(file, -1)
				if len(indices) == 1 {
					char := byte((strings.Count(string(file[:indices[0][0]]), "\n") / 9) + 33)
					res += string(char)
				}
			}
		}
		results = append(results, res)
	}
	azer := 0
	for _, v := range results {
		if len(v) > azer {
			azer = len(v)
		}
	}
	for _, v := range results {
		if len(v) == azer {
			fmt.Print(v)
		}
	}
}

// handleCharacter checks if the character can be part of the ASCII art
func handleCharacter(c rune) {
	if !strings.Contains("/_\\|,.'()` V<>", string(c)) {
		cannotBe["standard"] = false
	}
	if !strings.Contains(" _|", string(c)) {
		cannotBe["shadow"] = false
	}
	if !strings.Contains("o /O0-'|\\r\\", string(c)) {
		cannotBe["thinkertoy"] = false
	}
}
