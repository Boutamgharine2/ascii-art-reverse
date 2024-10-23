package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	asci "asci/ascii"
)

func main() {
	// fmt.Println(os.Args[1])
	if len(os.Args) == 2 {
		if !reverse(os.Args[1]) {
			asci.Simple(os.Args[1])
		}
	} else if len(os.Args) == 3 {
		asci.Fs(os.Args[1:])
	} else if len(os.Args) == 4 {
		asci.Complex(os.Args[1:])
	} else {
		fmt.Println("\n\t\t\t\t\t\t\tUsage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("\n\t\t\t\t\t\t\tor : go run . <text>")
		fmt.Println("\n\t\t\t\t\t\t\tEX: go run . --output=<fileName.txt> something standard")
		fmt.Print("\n\t\t\t\t\t\t\tEX: go run . \"Hello World\"\n\n")
		fmt.Println("\n\t\t\t\t\t\t\tBanners : standard, shadow, thinkertoy")
		fmt.Println("\n\t\t\t\t\t\t\tor : go run . --reverse=<fileName>")
	}
}

var cannbe map[string]bool

func reverse(s string) bool {
	cannbe = make(map[string]bool)
	cannbe["shadow"] = true
	cannbe["thinkertoy"] = true
	cannbe["standard"] = true
	fmt.Println("dkhol la reverse mode", s, strings.HasPrefix(s, "--reverse="))
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
	// res := []int{}
	counter := 0
	g := []bool{}
	for i, line := range data {
		if counter == 8 {
			counter = 0
		}

		if data[i] == "" {
			if counter != 0 {
				fmt.Printf("ERROR: Malformed file.\nLine %v should not be empty line.\n", i)
				os.Exit(0)
			}
			good = append(good, [8]string{})
			continue
		}
		counter++
		if counter == 1 || i == len(data)-1 {
			addCharacher(i+1, g, data)
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
	if counter == 0 {
		addCharacher(len(data)-1, g, data)
	}
	// fmt.Println(counter, good)
	resverseascii()
	return true
}

var good = [][8]string{}

func addCharacher(i int, g []bool, data []string) {
	fmt.Println("addCharacher", i, g)
	if i < 8 {
		return
	}
	i -= 9
	charachter := [8]string{}
	for i1, c := range g {
		if !c {
			for _, v := range charachter {
				fmt.Println(v)
			}
			fmt.Println("last char", charachter[6])
			fmt.Println("last char", charachter[7])
			good = append(good, charachter)
			charachter = [8]string{}
		} else {
			charachter[0] += string(data[0+i][i1])
			charachter[1] += string(data[1+i][i1])
			charachter[2] += string(data[2+i][i1])
			charachter[3] += string(data[3+i][i1])
			charachter[4] += string(data[4+i][i1])
			charachter[5] += string(data[5+i][i1])
			charachter[6] += string(data[6+i][i1])
			charachter[7] += string(data[7+i][i1])

		}
	}
	// good = append(good, [8]string{})
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
	res := ""
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(cannbe, otherthansp(""))
	for _, v := range good {
		rgx := "\n\n"
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
			res += " "
		}
	}
	fmt.Printf("'%v'\n", res)
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
