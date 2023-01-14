package main

var age int
var name string
var command, param1, param2, param3 string

type Player struct {
	name string
	age  int
	PlayerState
}

type Room struct {
	name       string
	accessRoom []string
	RoomState
}

type PlayerState struct {
	room string
	// artifactsPlayer []string
}

type RoomState struct {
	artifactsRoom map[string]bool
}

type fn func(string, string, string)