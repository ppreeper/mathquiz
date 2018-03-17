package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ppreeper/stringpad"
)

func main() {
	fmt.Println("Welcome to Math Quiz")
	question("add", 5)
}

func question(symbol string, level int) {
	firstNumber := 11111
	secondNumber := 11111
	answer := firstNumber * secondNumber
	firstString := stringpad.RJustLen(strconv.Itoa(firstNumber), level+2)
	secondString := stringpad.RJustLen("x "+strconv.Itoa(secondNumber), level+2)
	totalLine := stringpad.LPadLen("-", "-", level+2)
	fmt.Printf("%s\n", firstString)
	fmt.Printf("%s\n", secondString)
	fmt.Printf("%s\n", totalLine)
	totalString := stringpad.RJustLen(strconv.Itoa(firstNumber+secondNumber), level+2)

	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadString('\n')
	ansString := strings.TrimSpace(ans)
	ansInt, err := strconv.Atoi(ansString)
	if err != nil {
		panic(err)
	}
	if ansInt == answer {
		fmt.Printf("Correct\n")
	} else {
		fmt.Printf("Nope: %s\n", totalString)
	}

}
