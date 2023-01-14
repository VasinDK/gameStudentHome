package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func lookAround(par1 string, par2 string, par3 string) {
	fmt.Println("lookAround")
}

func move(par1 string, par2 string, par3 string) {
	fmt.Println("move")
}

func clothe(par1 string, par2 string, par3 string) {
	fmt.Println("clothe")
}

func take(par1 string, par2 string, par3 string) {
	fmt.Println("take")
}

func applay(par1 string, par2 string, par3 string) {
	fmt.Println("applay")
}

func initGame() {
	player = Player{
		"",
		0,
		PlayerState{startRoom, map[string]string{}},
	}

	fmt.Println("Игра началась. Введите команду или Help для правил")

}

func readCommand() {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Ошибка ввода, повторите попытку")
	}

	commandStr = str
}

func getCommandPlayer() {
	fmt.Println("Введите команду и 3 параметра")
}

func handleCommand(cmnd string) {
	command = strings.Fields(cmnd)
	fmt.Println(command[0])

	// fmt.Println(param2)
	// fmt.Println(param3)
	// проверить заполненость параметров тут перед вызовом
	// commandRun["осмотреться"]("1", "2", "3")
}
