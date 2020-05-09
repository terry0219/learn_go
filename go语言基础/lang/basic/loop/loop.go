package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// for 起始条件;结束条件;递增表达式
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

/*
读文件内容
file, err:= os.Open(filename)
scanner := bufio.NewScanner(file)
for scanner.Scan() {                 //scanner.Scan()这里是结束条件；这里省略了起始条件和递增表达式
	fmt.Println(scanner.Text())
}
*/

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//无限循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

/*
for的条件里不需要括号
for的条件里可以省略初始条件，结束条件，递增表达式
*/
func main() {
	fmt.Println("convertToBin results:")
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(72387885),
		convertToBin(0),
	)

	fmt.Println("abc.txt contents:")
	printFile("lang/basic/branch/abc.txt")

	fmt.Println("printing a string:")
	s := `abc"d"
	kkkk
	123

	p`
	printFileContents(strings.NewReader(s))

	// Uncomment to see it runs forever
	// forever()
}
