package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("What's your question ? \n")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "plus":
		// not using `sum` because it's a default command in unix.
		if len(arrCommandStr) < 3 {
			return errors.New("Required for 2 arguments")
		}
		arrNum := []int64{}
		for i, arg := range arrCommandStr {
			if i == 0 {
				continue
			}
			n, _ := strconv.ParseInt(arg, 10, 64)
			arrNum = append(arrNum, n)
		}
		fmt.Fprintln(os.Stdout, sum(arrNum...))
	// add another case here for custom commands.
	default:
		fmt.Print("I have no idea what you are talking about\n\n")
	}
	return nil
}

func sum(numbers ...int64) int64 {
	res := int64(0)
	for _, num := range numbers {
		res += num
	}
	return res
}
