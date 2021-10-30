package main

import (
	"math/rand"
	"time"

	sprite "github.com/pdevine/go-asciisprite"
	tm "github.com/pdevine/go-asciisprite/termbox"
)

var random *rand.Rand
var allSprites sprite.SpriteGroup
var Width int
var Height int

const cloudC0Timer = 22
const cloudC1Timer = 19

func NewCloud(cloudType, posX, posY int) *Cloud {
	c := &Cloud{BaseSprite: sprite.BaseSprite{
		Visible: true,
		Y:       posY,
		X:       posX},
		VX: -1,
	}
	if cloudType == 0 {
		c.AddCostume(sprite.NewCostume(cloudC0, 'x'))
		c.TimeOut = cloudC0Timer
	} else if cloudType == 1 {
		c.AddCostume(sprite.NewCostume(cloudC1, 'x'))
		c.TimeOut = cloudC1Timer
	}

	return c
}

func (c *Cloud) Update() {
	c.Timer++
	if c.Timer > c.TimeOut {
		c.X = c.X + c.VX

		if c.VX < 0 && c.X+c.Width < 0 {
			c.X = 102
		}
		c.Timer = 0
	}
}

func NewRip() *sprite.BaseSprite {
	c := &sprite.BaseSprite{
		Visible: true,
		X:       100,
		Y:       25,
		Alpha:   'x',
	}
	c.AddCostume(sprite.NewCostume(tombstoneAsset, 'x'))
	return c
}

func NewPumpkin(x, y int) *sprite.BaseSprite {
	c := &sprite.BaseSprite{
		Visible: true,
		X:       x,
		Y:       y,
		Alpha:   'x',
	}
	c.AddCostume(sprite.NewCostume(pumpkin_1, 'x'))
	return c
}

func NewHouse(x, y int) *sprite.BaseSprite {
	c := &sprite.BaseSprite{
		Visible: true,
		X:       x,
		Y:       y,
		Alpha:   'x',
	}
	c.AddCostume(sprite.NewCostume(houseAsset, 'x'))
	return c
}

func NewTitle(x, y int) *sprite.BaseSprite {
	c := &sprite.BaseSprite{
		Visible: true,
		X:       x,
		Y:       y,
		Alpha:   'x',
	}
	c.AddCostume(sprite.NewCostume(titleAsset, 'x'))
	return c
}

func NewGhost() *Ghost {
	g := &Ghost{BaseSprite: sprite.BaseSprite{
		Visible: true,
		X:       random.Intn(100) + 25,
		Y:       30},
		TimeOut: 3,
	}
	g.AddCostume(sprite.NewCostume(ghostC0, 'x'))
	g.AddCostume(sprite.NewCostume(ghostC1, 'x'))
	return g
}

func (g *Ghost) Update() {
	var x int
	var y int
	g.Timer++
	if g.Timer >= g.TimeOut {
		x = random.Intn(3)
		x--
		y = random.Intn(3)
		y--
		if x < 0 {
			g.CurrentCostume = 0
		} else if x > 0 {
			g.CurrentCostume = 1
		}
		g.X += x
		g.Y += y
		g.Timer = 0
	}
}

func NewBat() *Bat {
	b := &Bat{BaseSprite: sprite.BaseSprite{
		X:       5,
		Y:       5,
		Visible: true},
		VX: 1,
		VY: 1,
	}
	b.AddCostume(sprite.NewCostume(batAsset, 'x'))

	return b
}

func (bat *Bat) Update() {
	bat.X += bat.VX
	bat.Y += bat.VY
	if bat.X <= 0 {
		bat.VX = 1
	} else if bat.X >= (Width - bat.Width) {
		bat.VX = -1
	}
	if bat.Y <= 0 {
		bat.VY = 1
	} else if bat.Y >= (Height - bat.Height) {
		bat.VY = -1
	}
}

func main() {
	time.Sleep(1000 * time.Millisecond)

	err := tm.Init()
	if err != nil {
		panic(err)
	}
	defer tm.Close()

	Width, Height = tm.Size()

	random = rand.New(rand.NewSource(time.Now().UnixNano()))

	event_queue := make(chan tm.Event)
	go func() {
		for {
			event_queue <- tm.PollEvent()
		}
	}()

	txt := "Press 'ESC' to quit. 'Up' to spawn ghosts. Twitter: @_ediri"
	c := sprite.NewCostume(txt, '~')
	text := sprite.NewBaseSprite(Width/2-len(txt)/2, Height-2, c)

	cl0 := NewCloud(0, 202, 2)
	cl1 := NewCloud(1, 171, 2)
	cl2 := NewCloud(1, 123, 3)

	rip := NewRip()
	pumpkin := NewPumpkin(10, 35)
	pumpkin2 := NewPumpkin(65, 25)
	pumpkin3 := NewPumpkin(165, 35)
	hauntedHouse := NewHouse(5, -10)
	bat := NewBat()

	title := NewTitle(65, 0)

	for _, spr := range []sprite.Sprite{title, cl1, cl2, cl0, bat, rip, pumpkin, pumpkin2, pumpkin3, hauntedHouse, text} {
		allSprites.Sprites = append(allSprites.Sprites, spr)
	}

mainloop:
	for {
		tm.Clear(tm.ColorDefault, tm.ColorDefault)
		select {
		case ev := <-event_queue:
			if ev.Type == tm.EventKey {
				if ev.Key == tm.KeyCtrlC || ev.Key == tm.KeyEsc || ev.Ch == 'q' {
					break mainloop
				} else if ev.Key == tm.KeyArrowLeft {
					//	w.FaceLeft()
				} else if ev.Key == tm.KeyArrowRight {
					//	w.FaceRight()
				} else if ev.Key == tm.KeyArrowUp {
					g := NewGhost()
					allSprites.Sprites = append(allSprites.Sprites, g)
				}
			} else if ev.Type == tm.EventResize {
				Width = ev.Width
				Height = ev.Height
			}
		default:
			allSprites.Update()
			allSprites.Render()
			time.Sleep(50 * time.Millisecond)
		}
	}
}
