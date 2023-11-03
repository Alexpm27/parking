package models

import (
	"fyne.io/fyne/v2/canvas"
)

type Car struct {
	imgV  *canvas.Image
	imgH  *canvas.Image
	posX  int
	posY  int
	index int
}

func NewCar() *Car {
	return &Car{
		imgV: canvas.NewImageFromFile("img/lamboV.png"),
		imgH: canvas.NewImageFromFile("img/lamboH.png"),
		posX: 0,
		posY: 0,
	}
}

func (c *Car) ImgH() *canvas.Image {
	return c.imgH
}

func (c *Car) SetImgH(img *canvas.Image) {
	c.imgH = img
}

func (c *Car) ImgV() *canvas.Image {
	return c.imgV
}

func (c *Car) SetImgV(img *canvas.Image) {
	c.imgV = img
}

func (c *Car) SetIndex(inIndex int) {
	c.index = inIndex
}

func (c *Car) GetPosX() int {
	return c.posX
}

func (c *Car) SetPosX(pox int) {
	c.posX = pox
}

func (c *Car) GetPosY() int {
	return c.posY
}

func (c *Car) SetPosY(poy int) {
	c.posY = poy
}

func (c *Car) GetIndex() int {
	return c.index
}
