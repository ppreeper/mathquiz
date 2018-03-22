package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ppreeper/stringpad"
)

var randSource = rand.NewSource(time.Now().UnixNano())

func main() {
	fmt.Println("Welcome to Math Quiz")
	for i := 0; i < 5; i++ {
		question("add", 1)
	}
}

func question(symbol string, level int) {
	var fnum int
	var snum int
	var answer int
	var anslen int
	var firstString string
	var secondString string

	if level > 1 {
		fnum, _ = strconv.Atoi(stringpad.RPadLen("1", "0", level))
		snum, _ = strconv.Atoi(stringpad.LPadLen("", "9", level))
	} else {
		fnum, _ = strconv.Atoi(stringpad.RPadLen("0", "0", level))
		snum, _ = strconv.Atoi(stringpad.LPadLen("", "9", level))
	}

	var r = rand.New(randSource)
	firstNumber := r.Intn(snum-fnum) + fnum
	secondNumber := r.Intn(snum-fnum) + fnum
	switch symbol {
	case "add":
		answer = firstNumber + secondNumber
		anslen = len(strconv.Itoa(snum+snum)) + 2
		firstString = stringpad.RJustLen(strconv.Itoa(firstNumber), anslen)
		secondString = stringpad.RJustLen("+ "+strconv.Itoa(secondNumber), anslen)

	}
	totalLine := stringpad.LPadLen("-", "-", anslen)
	fmt.Printf("%s\n", firstString)
	fmt.Printf("%s\n", secondString)
	fmt.Printf("%s\n", totalLine)
	totalString := stringpad.RJustLen(strconv.Itoa(answer), anslen)

	reader := bufio.NewReader(os.Stdin)
	ans, _ := reader.ReadString('\n')
	ansString := strings.TrimSpace(ans)
	ansInt, err := strconv.Atoi(ansString)
	if err != nil {
		panic(err)
	}
	if ansInt == answer {
		fmt.Printf("\rCorrect\n")
	} else {
		fmt.Printf("\rNope: %s\n", totalString)
	}

}
