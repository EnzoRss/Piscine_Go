package main

import "z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func (d *Door) CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	d.state = "CLOSE"
}

func (d *Door) OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening")
	d.state = "OPEN"
}

func (d Door) IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?")
	return d.state == "OPEN"
}

func (d Door) IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return d.state == "CLOSE"
}

func main() {
	door := &Door{}

	door.OpenDoor(door)
	if door.IsDoorClose(door) {
		door.OpenDoor(door)
	}
	if door.IsDoorOpen(door) {
		door.CloseDoor(door)
	}
	if door.state == "OPEN" {
		door.CloseDoor(door)
	}
}
