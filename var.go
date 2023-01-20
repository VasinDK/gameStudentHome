package main

var player Player

var age int
var name string
var commandStr string
var command []string

var startRoom = "кухня"

var commandRun = map[string]fn{
	"осмотреться": lookAround,
	"идти":        move,
	"одеть":       clothe,
	"взять":       take,
	"применить":   applay,
}

var mapArtifacts = map[string]Artifact{
	"рюкзак": {
		"рюкзак",
		"пусто",
		"пусто",
	},
	"ключи": {
		"ключи",
		"рюкзак",
		"некуда класть",
	},
	"конспекты": {
		"конспекты",
		"пусто",
		"пусто",
	},
}

var backpack = Artifact{
	"рюкзак",
	"пусто",
	"пусто",
}

var keys = Artifact{
	"ключи",
	"рюкзак",
	"некуда класть",
}

var notes = Artifact{
	"конспекты",
	"пусто",
	"пусто",
}

var rooms = map[string]Room{
	"кухня": {
		accessRoom: map[string]bool{
			"коридор": true,
		},
		msgAfterMove:  "кухня, ничего интересного.",
		msgLookAround: "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ.",
		RoomState:     RoomState{map[string]map[string]Artifact{}},
	},
	"коридор": {
		accessRoom: map[string]bool{
			"кухня":   true,
			"комната": true,
			"улица":   true,
		},
		msgAfterMove:  "ничего интересного.",
		msgLookAround: "",
		RoomState:     RoomState{map[string]map[string]Artifact{}},
	},
	"комната": {
		accessRoom: map[string]bool{
			"коридор": true,
		},
		msgAfterMove:  "ты в своей комнате.",
		msgLookAround: "",
		RoomState: RoomState{
			map[string]map[string]Artifact{
				"на столе: ": {
					"ключи":     mapArtifacts["ключи"],
					"конспекты": mapArtifacts["конспекты"],
				},
				"на стуле - ": {
					"рюкзак": mapArtifacts["рюкзак"],
				},
			}},
	},
}
