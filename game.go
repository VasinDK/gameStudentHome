package main

import (
	"fmt"
)

var commandRun = map[string]fn{
	"осмотреться": lookAround,
	"идти":        move,
	"одеть":       clothe,
	"взять":       take,
	"применить":   applay,
}

func initGame() Player {
	player := Player{
		"",
		0,
		PlayerState{"кухня"},
	}

	// проверить заполненость параметров тут перед вызовом
	commandRun["осмотреться"]("1", "2", "3")

	return player
}

func getCommandPlayer() {
	fmt.Println("Введите команду и 3 параметра")
	fmt.Scan(&command, &param1, &param2, &param3)
}

func main() {
	player := initGame()
	fmt.Println(player)
	// fmt.Println("Введите имя и возрост")
	// fmt.Scan(&age, &name)
	// fmt.Println(age)
	// fmt.Println(name)
	// getCommandPlayer()
}
