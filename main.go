package main

import "./game"

/*
	컴퓨터와 한다고 가정
	컴퓨터가 문제를 낸다.
	사용자는 문제를 푼다.
	숫자는 3자리로 한다.
*/

func main() {
	baseball := game.Baseball{}
	baseball.Play()
}
