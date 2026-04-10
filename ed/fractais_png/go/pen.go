package main

import (
	"math"

	"github.com/fogleman/gg"
)

type Pen struct {
	x, y    float64
	angle   float64
	penDown bool
	dc      *gg.Context
}

// Cria uma nova instância de Pen com um canvas
func NewPen(width, height int) *Pen {
	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(2)
	return &Pen{
		x: float64(width) / 2, y: float64(height) / 2,
		angle: 0, penDown: true, dc: dc,
	}
}

// Move a caneta dist unidades na direção atual.
func (t *Pen) Walk(dist float64) {
	newX := t.x + dist*math.Cos(t.angle*math.Pi/180)
	newY := t.y - dist*math.Sin(t.angle*math.Pi/180)
	if t.penDown {
		t.dc.DrawLine(t.x, t.y, newX, newY)
		t.dc.Stroke()
	}
	t.x, t.y = newX, newY
}

// Gira a direção da caneta deg graus no sentido anti-horário
func (t *Pen) Left(deg float64) { t.angle += deg }

// Gira a direção da caneta deg graus no sentido horário (diminui o ângulo).
func (t *Pen) Right(deg float64) { t.angle -= deg }

// Levanta a caneta.
func (t *Pen) Up() { t.penDown = false }

// Abaixa a caneta.
func (t *Pen) Down() { t.penDown = true }

// Move a caneta diretamente para as coordenadas
func (t *Pen) Goto(x, y float64) {
	if t.penDown {
		t.dc.DrawLine(t.x, t.y, x, y)
		t.dc.Stroke()
	}
	t.x, t.y = x, y
}

// Move a caneta para as coordenadas (x, y) sem desenhar
func (t *Pen) SetPosition(x, y float64) {
	t.x = x
	t.y = y
}

// Define a direção atual da caneta
func (t *Pen) SetHeading(angle float64) {
	t.angle = angle
}

// Desenha o contorno de um círculo com o centro na posição atual e raio
func (t *Pen) DrawCircle(radius float64) {
	if t.penDown {
		t.dc.DrawCircle(t.x, t.y, radius)
		t.dc.Stroke()
	}
}

// Desenha o contorno de um retângulo com largura w e altura h
func (t *Pen) DrawRect(w, h float64) {
	if t.penDown {
		t.dc.DrawRectangle(t.x, t.y, w, h)
		t.dc.Stroke()
	}
}

// Desenha e preenche um círculo com raio
func (t *Pen) FillCircle(radius float64) {
	if t.penDown {
		t.dc.DrawCircle(t.x, t.y, radius)
		t.dc.Fill()
	}
}

// Desenha e preenche um retângulo (ou quadrado) com largura w e altura h
func (t *Pen) FillSquare(w, h float64) {
	t.dc.DrawRectangle(t.x, t.y, w, h)
	t.dc.Fill()
}

// Define a cor da caneta com os valores RGB fornecidos
func (t *Pen) SetRGB(r, g, b float64) {
	t.dc.SetRGB(r/255, g/255, b/255)
}

// Define a espessura da linha usada para desenhar.
func (t *Pen) SetLineWidth(w float64) {
	t.dc.SetLineWidth(w)
}

// Salva a imagem atual em um arquivo PNG com o nome e caminho fornecido.
func (t *Pen) SavePNG(path string) {
	t.dc.SavePNG(path)
}
