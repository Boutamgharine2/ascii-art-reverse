package main_test

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	ascii "azer/ascii"
)

func TestAsci(t *testing.T) {
	for lib, test := range tests {
		for v, v1 := range test {
			o := strings.Join(ascii.Asci(v, lib), "")
			f, err := os.ReadFile(v1)
			if err != nil {
				log.Fatal(err)
			}
			g := string(f)
			if o != g {
				t.Fail()
				log.Printf("failed the %v test whith the %s \noutput:\n[%v]", v1, lib, o)
			}
		}
	}
}

var tests = map[string]map[string]string{
	"standard.txt":   ts,
	"shadow.txt":     tshad,
	"thinkertoy.txt": tthink,
}

var ts = map[string]string{
	"t":                                     "tests/t_test",
	"G":                                     "tests/G_test",
	"Hello World":                           "tests/Hello_World_test",
	"1":                                     "tests/1_test",
	"233546612":                             "tests/233546612_test",
	"123456789":                             "tests/123456789_test",
	"1234567890":                            "tests/1234567890_test",
	"abcdefghijklmopqr":                     "tests/abcdefghijklmopqr_test",
	"ABCDEFGHIJKLMOPQRTUVWXYZ":              "tests/ABCDEFGHIJKLMOPQRTUVWXYZ_test",
	"testereffsefdf":                        "tests/testereffsefdf_test",
	"01234567890":                           "tests/01234567890_test",
	"\\n\\n\\n\\n\\ngogog power rangers\\n": "tests/gogo power rangers_test",
	"\\naaaaaaaa":                           "tests/aaaaaaaa1_test",
	"\\n\\naaaaaaaa":                        "tests/aaaaaaaa2_test",
	"\\n\\n\\naaaaaaaa":                     "tests/aaaaaaaa3_test",
	"bbbbbbbb\\n":                           "tests/bbbbbbbb1_test",
	"bbbbbbbb\\n\\n":                        "tests/bbbbbbbb2_test",
	"bbbbbbbb\\n\\n\\n":                     "tests/bbbbbbbb3_test",
	"\\ncccccccc\\n":                        "tests/cccccccc1_test",
	"\\n\\ncccccccc\\n\\n":                  "tests/cccccccc2_test",
	"\\n\\n\\ncccccccc\\n\\n\\n":            "tests/cccccccc3_test",
	"\\n\\n\\namericano\\n":                 "tests/americano_test",
	"\\n\\n":                                "tests/emty_test",
	" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~": "tests/all_test",
}

var tshad = map[string]string{
	"t":                                     "tests/t_test_shadow",
	"G":                                     "tests/G_test_shadow",
	"Hello World":                           "tests/Hello_World_test_shadow",
	"1":                                     "tests/1_test_shadow",
	"233546612":                             "tests/233546612_test_shadow",
	"123456789":                             "tests/123456789_test_shadow",
	"1234567890":                            "tests/1234567890_test_shadow",
	"abcdefghijklmopqr":                     "tests/abcdefghijklmopqr_test_shadow",
	"ABCDEFGHIJKLMOPQRTUVWXYZ":              "tests/ABCDEFGHIJKLMOPQRTUVWXYZ_test_shadow",
	"testereffsefdf":                        "tests/testereffsefdf_test_shadow",
	"01234567890":                           "tests/01234567890_test_shadow",
	"\\n\\n\\n\\n\\ngogog power rangers\\n": "tests/gogo power rangers_test_shadow",
	"\\n\\n\\namericano\\n":                 "tests/americano_test_shadow",
	" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~": "tests/all_test_shadow",
}

var tthink = map[string]string{
	"t":                                     "tests/t_test_think",
	"G":                                     "tests/G_test_think",
	"Hello World":                           "tests/Hello_World_test_think",
	"1":                                     "tests/1_test_think",
	"233546612":                             "tests/233546612_test_think",
	"123456789":                             "tests/123456789_test_think",
	"1234567890":                            "tests/1234567890_test_think",
	"abcdefghijklmopqr":                     "tests/abcdefghijklmopqr_test_think",
	"ABCDEFGHIJKLMOPQRTUVWXYZ":              "tests/ABCDEFGHIJKLMOPQRTUVWXYZ_test_think",
	"testereffsefdf":                        "tests/testereffsefdf_test_think",
	"01234567890":                           "tests/01234567890_test_think",
	"\\n\\n\\n\\n\\ngogog power rangers\\n": "tests/gogo power rangers_test_think",
	"\\n\\n\\namericano\\n":                 "tests/americano_test_think",
	" !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~": "tests/all_test_think",
}

func TestReverse(t *testing.T) {
	for _, testgrp := range tests {
		for i, v := range testgrp {
			out, err := exec.Command("go", "run", ".", "--reverse="+v).CombinedOutput()
			if err != nil {
				t.Error(v, string(out), err)
			}
			if string(out) != i {
				t.Fail()
				log.Printf("failed the %v test,'%v',good '%v'", v, string(out), i)
			}
		}
	}
}
