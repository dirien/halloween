package main

import sprite "github.com/pdevine/go-asciisprite"

type Cloud struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
	VX      int
}

type Ghost struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
}

type Bat struct {
	sprite.BaseSprite
	VX int
	VY int
}
