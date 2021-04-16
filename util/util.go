package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func GetInput() string {
	fmt.Println("각 자리가 다른 3자리 수를 입력하세요.")
	scanner.Scan()
	answer := scanner.Text()

	_, err := strconv.Atoi(answer)
	if err != nil {
		fmt.Println("숫자를 입력해 주세요.")
		return ""
	}

	if len(answer) != 3 {
		fmt.Println("3자리 수를 입력해 주세요.")
		return ""
	}

	if answer[0] == answer[1] || answer[1] == answer[2] || answer[2] == answer[0] {
		fmt.Println("각 자리가 다른 수를 입력하세요.")
		return ""
	}

	return answer
}
