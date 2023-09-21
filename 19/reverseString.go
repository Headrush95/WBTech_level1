package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

/*
 Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func readWord(in io.Reader) string {
	rdr := bufio.NewReader(in)
	res, err := rdr.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Trim(res, "\n\r")
}

// reverseString разбивает строку на руны и поочередно меняет местами
func reverseString(src string) string {
	res := []rune(src)
	length := utf8.RuneCountInString(src) - 1
	for i := 0; i < length/2+length%2; i++ {
		res[i], res[length-i] = res[length-i], res[i]
	}
	return string(res)
}

func main() {
	testString := readWord(os.Stdin)
	fmt.Println(reverseString(testString))
}
