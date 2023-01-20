package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getMessage(room Room, firstMsg string) string {
	msg := firstMsg
	msg += " можно пройти - "

	accessRoomCount := len(room.accessRoom)
	index := 0
	for key, _ := range room.accessRoom {
		msg += key

		if index < accessRoomCount-1 {
			msg += ", "
			index++
		}
	}

	return msg
}

func noWay(way string) string {
	msg := "нет пути в "
	msg += way

	return msg
}

func deleteItemSl(sl *[]string, i int) {
	(*sl)[i] = (*sl)[len((*sl))-1]
	(*sl) = (*sl)[:len((*sl))-1]
}

func initGame() {
	player = Player{
		"",
		0,
		PlayerState{startRoom, map[string]Artifact{}},
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

func handleCommand(cmnd string) string {
	readCommand()

	msg := ""
	command = strings.Fields(cmnd)

	if len(command) < 4 {
		command = append(command, "", "", "")
	}

	run, ok := commandRun[command[0]]

	if !ok {
		msg = "неизвестная команда"
	}

	if ok {
		msg = run(command[1], command[2], command[3])
	}

	fmt.Println(msg)

	// handleCommand(commandStr)
	return msg

	// во всех остальных случаях когда нет command[0] добавить
}
