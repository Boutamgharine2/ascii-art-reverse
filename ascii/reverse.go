package asci

import (
	"fmt"
	"log"
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
	counter := 0
	var g []bool

	for i, line := range data {
		if counter == 8 {
			counter = 0
		}

		if counter == 0 && i == len(data)-1 {
			good = append(good, addCharacter(g, [8]string(data[i-8:i]))...)
			break
		}

		if line == "" && i < len(data)-1 {
			if counter != 0 {
				fmt.Printf("ERROR: Malformed file.\nLine %v should not be an empty line.\n", i)
				os.Exit(0)
			}
			good = append(good, [8]string{"\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n"})
			continue
		}

		counter++
		if counter == 1 || i == len(data)-1 {
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
				character[j] += string(data[j][i1])
			}
		}
	}
	lines = append(lines, [8]string{"\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n"})
	return lines
}

// otherThanSpace checks if the string contains any non-space character
func otherThanSpace(a string) bool {
	for _, v := range a {
		if v != ' ' {
			return true
		}
	}
	return false
}

// reverseASCII processes the ASCII art to generate the final output
func reverseASCII() {
	if len(good) == 0 {
		fmt.Println(0)
		return
	}
	good = good[:len(good)-1]
	res := ""
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	counter := 0
	for _, v := range good {
		rgx := "\n\n"
		if v[0] == " " {
			counter++
			if counter == 6 {
				res += " "
				counter = 0
			}
			continue
		} else if v[0] == "\n" {
			res += "\\n"
			continue
		}

		for _, v1 := range v {
			if !otherThanSpace(v1) {
				continue
			}
			rgx += "\\s*" + regexp.QuoteMeta(v1) + "\\s+\n"
		}

		if len(rgx) > 2 {
			regex := regexp.MustCompile(rgx)
			indices := regex.FindAllIndex(file, -1)
			if len(indices) == 1 {
				char := byte((strings.Count(string(file[:indices[0][0]]), "\n") / 9) + 33)
				res += string(char)
			}
		} else {
			log.Fatalln(v)
		}
	}
	fmt.Printf("%v", res)
}

// handleCharacter checks if the character can be part of the ASCII art
func handleCharacter(c rune) {
	if !strings.Contains("/_\\|,.'()` ", string(c)) {
		cannotBe["standard"] = false
	}
	if !strings.Contains(" _|", string(c)) {
		cannotBe["shadow"] = false
	}
	if !strings.Contains("o /O-'|\\r\\", string(c)) {
		cannotBe["thinkertoy"] = false
	}
}
