package main

import (
	"sort"
)



func (p Player) lookAround() string {
	msg := ""
	artifactLocationCount := len(rooms[p.room].artifactsRoom)

	if artifactLocationCount > 0 {
		indexLocation := 0
		artifactsLocation := make([]string, 0, artifactLocationCount)

		for i, _ := range rooms[p.room].artifactsRoom {
			artifactsLocation = append(artifactsLocation, i)
		}

		sort.Strings(artifactsLocation)

		for _, v := range artifactsLocation {
			msg += v
			artifactCount := len(rooms[p.room].artifactsRoom[v])

			indexAtrifact := 0
			for i, _ := range rooms[p.room].artifactsRoom[v] {
				msg += i

				if indexAtrifact < artifactCount-1 {
					msg += ", "
				}
				indexAtrifact++
			}

			if indexLocation < artifactLocationCount-1 {
				msg += ", "
			}

			if indexLocation == artifactLocationCount-1 {
				msg += "."
			}

			indexLocation++
		}

	}

	msg += getMessage(rooms[player.room], rooms[player.room].msgLookAround)

	return msg
}

func (p *Player) move(par1 string) string {
	if rooms[p.room].accessRoom[par1] {
		p.room = par1

		return getMessage(rooms[p.room], rooms[p.room].msgAfterMove)
	}
	return noWay(par1)
}

func (p *Player) clothe(par1 string) string {
	msg := "нет такого"

	val, ok := rooms[p.room].artifactsRoom["на стуле - "]

	if !ok {
		return "Здесь нечего одевать"
	}

	for i, _ := range val {
		if par1 == i {
			p.artifactsPlayer[par1] = rooms[p.room].artifactsRoom["на стуле - "][par1]
			delete(rooms[p.room].artifactsRoom["на стуле - "], par1)
			msg = "вы одели: "
			msg += par1
			break
		}
	}

	return msg
}

func (p *Player) take(par1 string) string {
	msg := "нет такого"
	var take = true

	artifact, ok := rooms[p.room].artifactsRoom["на столе: "][par1]

	if !ok {
		return msg
	}

	if artifact.termsOfUse != "пусто" {
		_, take = p.artifactsPlayer[artifact.termsOfUse]
	}

	if take {
		delete(rooms[p.room].artifactsRoom["на столе: "], par1)
		p.artifactsPlayer[par1] = artifact

		msg = "предмет добавлен в инвентарь: "
		msg += par1
	}

	if !take {
		msg = artifact.failTakeMessage
	}

	return msg
}

