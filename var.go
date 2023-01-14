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
