package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func readWords(in io.Reader) string {
	rdr := bufio.NewReader(in)
	res, err := rdr.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	return strings.Trim(res, "\n\r")
}

// reverseString разбивает строку на слова и переворачивает их
func reverseWords(src string) string {
	res := strings.Fields(src)
	length := len(res) - 1
	for i := 0; i < length/2+length%2; i++ {
		res[i], res[length-i] = res[length-i], res[i]
	}
	return strings.Join(res, " ")
}
func main() {
	testString := readWords(os.Stdin)
	fmt.Println(reverseWords(testString))

}
