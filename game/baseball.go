package game

import (
	"fmt"
	"math/rand"
	"strings"

	"../util"
)

// 틀린 경우 결과 값 저장
type Result struct {
	ball   int
	strike int
	out    int
}

/*
	게임 진행
	정답 초기화
	입력 받기
	정답 체크
	결과 주기
*/
type Baseball struct {
	answer string
}

func (b *Baseball) Play() {
	b.initAnswer()

	for {
		input := util.GetInput()
		if input == "" {
			continue
		}

		if b.answer == input {
			fmt.Println("정답입니다.")
			return
		}

		fmt.Println(b.getResult(input))
	}
}

func (b *Baseball) initAnswer() {
	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < 3; i++ {
		index := rand.Intn(10 - i)
		b.answer += numbers[index]
		numbers = append(numbers[:index], numbers[index+1:]...)
	}
}

func (b *Baseball) getResult(input string) string {
	result := Result{}
	for i := 0; i < 3; i++ {
		index := strings.IndexByte(b.answer, input[i])
		if index == -1 {
			result.out += 1
		} else if index == i {
			result.strike += 1
		} else {
			result.ball += 1
		}
	}

	return fmt.Sprintf("%d 스트라이크, %d 볼, %d 아웃", result.strike, result.ball, result.out)
}
