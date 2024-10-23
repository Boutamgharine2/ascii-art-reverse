package asci

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Simple(in string) {
	fmt.Print(strings.Join(Asci(in, "standard.txt"), ""))
}

func Fs(in []string) {
	if len(f[in[1]]) == 0 {
		log.Fatal("\n\t\t\t\t\t\t\tNON VALID BANNER\n\nvalid banners : \"standard\" \"thinkertoy\" \"shadow\"\n")
	}
	fmt.Print(strings.Join(Asci(in[0], f[in[1]]), ""))
}

var (
	n   = ""
	nxe = ""
)

func Complex(in []string) {
	c := 0
	cu := ""
	for v := range q {
		err := FlagCheck(v, &in[0])
		if err != nil {
			c++
		} else {
			cu = v
			break
		}
		if c == 3 {
			log.Fatal("\n\t\t\t\t\t\t\tUsage: go run . [OPTION] [STRING] [BANNER]\n\t\t\t\t\t\t\tEX: go run . --output= <fileName.txt> something standard")
		}
	}

	if len(f[in[2]]) == 0 {
		log.Fatal("\n\t\t\t\t\t\t\tNON VALID BANNER\n\nvalid banners : \"standard\" \"thinkertoy\" \"shadow\"\n")
	}

	n, nxe = in[1], in[2]
	l := strings.Count(in[1], "\\n")
	if l*2 != len(in[1]) {
		o := Asci(in[1], f[in[2]])
		Format(&o, in[0], cu)
		// Write(o, in[0])
	} else {
		for l > 0 {
			fmt.Println()
			l--
		}
	}
}

func Write(o0 []string, f string) {
	o := []byte(strings.Join(o0, ""))
	var Error error
	var file *os.File
	if _, Error = os.Stat(f); os.IsNotExist(Error) {
		file, Error = os.Create(f)
		if Error != nil {
			panic(Error)
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		panic(Error)
	}
	if f == "main.go" {
		log.Fatal("output cannot be  name of the program")
	}
	Error = os.WriteFile(f, o, 0o644)
	if Error != nil {
		panic(Error)
	}
	fmt.Println("completed successfully")
}

func FlagCheck(v string, in *string) error {
	c := strings.Index(*in, v)
	c1 := strings.Count(*in, v)
	if c != 0 && c1 != 0 {
		return fmt.Errorf("not at 0")
	} else if c == 0 && c1 != 0 {
		*in = strings.Replace(*in, v, "", 1)
	} else if c == -1 {
		return fmt.Errorf("NotFound")
	}

	return nil
}

func Format(o *[]string, in string, v string) {
	if v == "--output=" {
		q[v](*o, in)
		return
	}
	if v == "--color=" {
		q[v](*o, in)
		return
	}
	if v == "--allign=" {
		q[v](*o, in)
		return
	}
	if v == "--reverse=" {
	}
}

func Color(o0 []string, v string) {
	o := []byte(strings.Join(o0, ""))
	if len(t[v]) == 0 {
		log.Fatal("\n\t\t\t\t\t\t\tNON VALID COLOR\n\nvalid colors : \"Black\" \"Red\" \"Green\" \"Yellow\" \"Blue\" \"Magenta\" \"Cyan\" \"White\"..........\n")
	}
	fmt.Print(t[v] + string(o) + t["ColorOff"])
}

func Allign(o []string, align string) {
	// Get terminal width
	if align != "center" && align != "right" && align != "left" && align != "justify" {
		fmt.Print(os.Args[1])
		log.Fatal("\ninvalid allighnment correct allignment is \nleft\nright\ncenter\njustify\nEX:go run . --allign=right something standard")
	}
	width, err := terminalWidth()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, v := range o {
		// message := "Hello, Hello, World!Hello World!, World!World!Hello World!, World!SHDK"
		message := v
		if len(v) > width {
			log.Fatal("too many characters in one line")
		}
		// Calculate padding based on alignment type
		var padding string
		switch align {
		case "center":
			padding = strings.Repeat(" ", (width-len(message))/2)
		case "right":
			padding = strings.Repeat(" ", width-len(message))
		case "justify":
			justifyText(width)
			return
		default: // Left alignment
			padding = ""
		}

		// Print the message with padding
		fmt.Print(padding + message)
	}
}

// terminalWidth returns the width of the terminal.
func terminalWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	size := strings.Fields(string(out))
	width, err := strconv.Atoi(size[1])
	if err != nil {
		return 0, err
	}
	return width, nil
}

// justifyText justifies the given text to fit the specified width.
func justifyText(width int) {
	text := n
	W_0 := strings.Split(text, "\\n")
	for _, v := range W_0 {
		if !multipleword(v) {
			Allign(Asci(v, f[nxe]), "center")
			continue
		}
		v = strings.TrimRight(v, " ")
		v2 := strings.Split(v, " ")
		v3 := [8]string{}
		for i1, v1 := range v2 {
			if v1 != "" && i1 != len(v2)-1 {
				v0 := Asci(v1, f[nxe])
				for i, x := range v0 {
					v3[i] += x + "\n"
				}
			}
			if v1 != "" && i1 == len(v2)-1 {
				v0 := Asci(v1, f[nxe])
				for i, x := range v0 {
					v3[i] += x[:len(x)-1]
				}
			}
		}
		for _, v := range v3 {
			ws := strings.Split(v, "\n")
			nSp := width - len(v) + len(ws) - 1
			sp := nSp / (len(ws) - 1)
			extraSp := nSp % (len(ws) - 1)

			var justT strings.Builder
			for i, word := range ws {
				justT.WriteString(word)
				if i < len(ws)-1 {
					justT.WriteString(strings.Repeat(" ", sp))
					if i < extraSp {
						justT.WriteByte(' ')
					}
				}
			}
			fmt.Print(justT.String())
		}
		fmt.Println()
	}
}

func multipleword(s string) bool {
	s = strings.Trim(s, " ")
	s1 := strings.Count(s, " ")
	if s1 == 0 {
		return false
	}
	return true
}
