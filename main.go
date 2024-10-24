package main

import (
	"fmt"
	"os"

	asci "azer/ascii"
)

func main() {
	// fmt.Println(os.Args[1])
	if len(os.Args) == 2 {
		if !asci.Reverse(os.Args[1]) {
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
