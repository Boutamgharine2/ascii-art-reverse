package asci

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	cannbe map[string]bool
	good   = [][8]string{}
)

func Reverse(s string) bool {
	cannbe = make(map[string]bool)
	cannbe["shadow"] = true
	cannbe["thinkertoy"] = true
	cannbe["standard"] = true
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
	// data = data[:len(data)-1]
	counter := 0
	g := []bool{}
	for i, line := range data {
		if counter == 8 {
			counter = 0
		}
		// fmt.Println(i, counter)
		if counter == 0 && i == len(data)-1 {
			good = append(good, addCharacher(g, [8]string(data[i-8:i]))...)
			break
		}
		if line == "" && i < len(data)-1 {
			// fmt.Println("i:====", i, good)
			if counter != 0 {
				fmt.Printf("ERROR: Malformed file.\nLine %v should not be empty line.\n", i)
				os.Exit(0)
			}
			good = append(good, [8]string{"\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n"})
			continue
		}
		counter++
		if counter == 1 || i == len(data)-1 {
			if i > 7 {
				good = append(good, addCharacher(g, [8]string(data[i-8:i]))...)
			}
			g = make([]bool, len(line))
		}
		// fmt.Println("counter", counter, i, line, len(g), len(line))
		if len(g) != len(line) {
			fmt.Printf("ERROR: Malformed file.\nLine %v should be %v charachers long.\n", i, len(g))
			os.Exit(0)
		}
		for i1, c := range line {
			if c != ' ' {
				g[i1] = true
				canntbehandler(c)
				continue
			}
		}

	}

	resverseascii()
	return true
}

func addCharacher(g []bool, data [8]string) [][8]string {
	// fmt.Println("data:", data) // Check if the value is false
	charachter := [8]string{}
	line := [][8]string{}
	for i1, c := range g {
		if !c {
			if charachter[0] != "" {
				line = append(line, charachter)
			} else {
				line = append(line, [8]string{" ", "", "", "", "", "", "", ""})
			}
			charachter = [8]string{}
		} else {
			charachter[0] += string(data[0][i1])
			charachter[1] += string(data[1][i1])
			charachter[2] += string(data[2][i1])
			charachter[3] += string(data[3][i1])
			charachter[4] += string(data[4][i1])
			charachter[5] += string(data[5][i1])
			charachter[6] += string(data[6][i1])
			charachter[7] += string(data[7][i1])
		}
	}
	line = append(line, [8]string{"\n", "\n", "\n", "\n", "\n", "\n", "\n", "\n"})
	return line
}

func otherthansp(a string) bool {
	for _, v := range a {
		if v != ' ' { // Check if the value is false
			return true
		}
	}
	return false
}

func resverseascii() {
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
	// fmt.Println(cannbe, otherthansp(""))
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
			if !otherthansp(v1) {
				continue
			}
			rgx += "\\s*" + regexp.QuoteMeta(v1) + "\\s+\n"
		}
		if len(rgx) != 2 {

			regex := regexp.MustCompile(rgx)
			azer := regex.FindAllIndex(file, -1)
			if len(azer) == 1 {

				// fmt.Println(azer)
				char := byte((strings.Count(string(file[:azer[0][0]]), "\n") / 9) + 33)
				// fmt.Println(char)
				res += string(char)
			} else {
				// fmt.Println(azer)
				// for _, qsdf := range v {
				// 	// fmt.Println(qsdf)
				// }
				// fmt.Println(rgx)
			}
		} else {
			log.Fatalln(v)
		}
	}
	fmt.Printf("%v", res)
}

func canntbehandler(c rune) {
	if !strings.Contains("/_\\|,.'()` ", string(c)) {
		// fmt.Println(string(c))
		cannbe["standard"] = false
	}
	if !strings.Contains(" _|", string(c)) {
		// fmt.Println(string(c))
		cannbe["shadow"] = false
	}
	if !strings.Contains("o /O-'|\r\\", string(c)) {
		// fmt.Println(string(c))
		cannbe["thinkertoy"] = false
	}
}

// func Reverse1(s string) bool {
// 	cannbe = make(map[string]bool)
// 	if !strings.HasPrefix(s, "--reverse=") {
// 		return false
// 	}
// 	s = strings.TrimPrefix(s, "--reverse=")
// 	file, err := os.ReadFile(s)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	// data = data[:len(data)-1]
// 	// res := ""
// 	// template := ParseTmplate("standard.txt")
// 	lines := splitdata(&file)
// 	fmt.Printf("%v\n", lines)
// 	return true
// }

// func splitdata(file *[]byte) [][]string {
// 	data := strings.Split(string(*file), "\n")
// 	var res [][]string
// 	line := []string{}
// 	for _, ln := range data {
// 		if len(ln) == 0 {
// 			res = append(res, []string{})
// 			continue
// 		}
// 		if len(line)%8 == 1 {
// 			res = append(res, []string{})
// 			continue
// 		} else if len(line)%8 != 0 {
// 			fmt.Println(len(line), len(v), len(line)%8)
// 			return nil
// 		}
// 		res = append(res, splitIntoChunks(line, 8)...)
// 	}
// 	fmt.Println(res, data)
// 	return res
// }

// func splitIntoChunks(data []string, chunkSize int) [][]string {
// 	var chunks [][]string
// 	for i := 0; i < len(data); i += chunkSize {
// 		end := i + chunkSize
// 		if end > len(data) {
// 			end = len(data)
// 		}
// 		chunks = append(chunks, data[i:end])
// 	}
// 	return chunks
// }
