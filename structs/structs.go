package main

import "fmt"

type postition struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    postition
}

func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}
func main() {

	p := postition{4, 2}
	fmt.Println(p.x)

	b := badGuy{"Jabba The Hut", 100, p}
	whereIsBadGuy(b)
}
