package asci

import (
	"log"
	"os"
	"strings"
)

var t1 = [95][8]string{}

func Asci(want string, f string) []string {
	k := strings.Count(want, "\\n")
	if k*2 == len(want) {
		return []string{strings.Replace(want, "\\n", "\n", -1)}
	}
	t1 := [95][8]string{}
	c, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err, " os.ReadFile")
	}
	c2 := []string{}
	if f == "thinkertoy.txt" {
		c2 = strings.Split(string(c), "\r\n")
	} else {
		c2 = strings.Split(string(c), "\n")
	}
	counter := 0
	for i, v := range t1 {
		counter++
		for j := range v {
			t1[i][j] = c2[counter]
			counter++
		}
	}
	return maker(want, &t1)
}

func maker(want string, t *[95][8]string) []string {
	ans := []string{}
	a := strings.Split(want, "\\n")
	for _, v := range a {
		ans = append(ans, build(v, t)...)
	}
	return ans
}

func build(want string, t *[95][8]string) []string {
	if len(want) == 0 {
		return []string{"\n"}
	}
	u := []byte(want)
	ans := make([]string, 8)
	for i := 0; i < 8; i++ {
		for y, v := range u {
			if v >= 32 && v <= 126 {
				ans[i] += (*t)[v-32][i]
			} else {
				log.Fatalf("cannot Print the charachter at index: %d", y)
			}
		}
		ans[i] += "\n"
	}
	return ans
}
