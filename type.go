package main

type Player struct {
	name string
	age  int
	PlayerState
}

type Room struct {
	accessRoom    map[string]bool
	msgAfterMove  string
	msgLookAround string
	RoomState
}

type PlayerState struct {
	room            string
	artifactsPlayer map[string]Artifact
}

type RoomState struct {
	artifactsRoom map[string]map[string]Artifact
}

type Artifact struct {
	name		string
	termsOfUse	string
	failTakeMessage string
}

type fn func(string, string, string) string
