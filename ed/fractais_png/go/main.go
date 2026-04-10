package main

import (
	"fmt"
	"math/rand"
)

func pentagonDrawing(p *Pen, tam float64) {
	for i := 0; i < 5; i++ {
		p.Walk(tam)
		p.Left(72)
	}

}

func fractalPentagon(p *Pen, tam float64, level int) {
	if level == 0 || tam < 5 {
		return
	}
	pentagonDrawing(p, tam)

	x1, y1 := p.x, p.y
	ang1 := p.angle

	for i := 0; i < 5; i++ {
		p.Up()
		p.Walk(tam)
		p.Down()

		fractalPentagon(p, tam*45, level-1)

		p.Up()
		p.SetPosition(x1, y1)
		p.SetHeading(ang1)
		p.Down()

		p.Left(72)
	}
}
func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	pen := NewPen(800, 800) // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(0, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetLineWidth(2)
	pen.SetPosition(400, 400)
	pen.SetHeading(0)

	for i := 0; i < 6; i++ {
		pen.Up()
		pen.Walk(200)
		pen.Down()
		fractalPentagon(pen, 60, 3)

		pen.Up()
		pen.SetPosition(400, 400)
		pen.SetHeading(float64(i * 60))
		pen.Down()
	}

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
