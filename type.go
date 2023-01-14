package main

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
	artifactsPlayer map[string]string
}

type RoomState struct {
	artifactsRoom map[string]bool
}

type fn func(string, string, string)