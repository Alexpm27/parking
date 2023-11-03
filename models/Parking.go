package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"sync"
	"time"
)

type Parking struct {
	capacity int
	occupied int
	mutex    sync.Mutex
	place    []bool
	posX     []int
	posY     []int
}

func NewParking(capacity int) *Parking {
	return &Parking{
		capacity: capacity,
		place:    make([]bool, capacity),
		posX: []int{
			205,
			302,
			400,
			495,
			590,
			685,
			785,
			880,
			975,
			1070,
			180,
			275,
			375,
			470,
			567,
			663,
			760,
			860,
			960,
			1055,
		},
		posY: []int{
			100,
			580,
		},
	}
}

func (p *Parking) ManageSpace(car *Car, carImgV *canvas.Image, carImgH *canvas.Image, sceneContainer *fyne.Container) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	time.Sleep(time.Second)
	for {
		if p.occupied < p.capacity {
			break
		}
		sceneContainer.Remove(carImgV)
		sceneContainer.Refresh()
		return
	}

	for i := 0; i < p.capacity; i++ {
		if !p.place[i] {
			p.place[i] = true
			p.occupied++

			if i < 10 {
				p.left(car, carImgH)
				carImgV.Move(fyne.NewPos(float32(p.posX[i]), float32(p.posY[0])))
				car.SetIndex(i)
				break
			}
			p.left(car, carImgH)
			carImgV.Move(fyne.NewPos(float32(p.posX[i]), float32(p.posY[1])))
			car.SetIndex(i)
			break
		}
	}
}

func (p *Parking) Exit(car *Car, carImgV *canvas.Image, carImgH *canvas.Image, sceneContainer *fyne.Container) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for i := p.capacity - 1; i >= 0; i-- {
		if i == car.index {
			p.place[i] = false
			p.occupied--

			if i < 10 {
				p.reverseAbove(carImgV, i)
				sceneContainer.Remove(carImgV)
				sceneContainer.Refresh()
				break
			}

			p.reverseAbove(carImgV, i)
			sceneContainer.Remove(carImgV)
			sceneContainer.Refresh()
			break
		}
	}
}

//animations

func (p *Parking) left(car *Car, carImgH *canvas.Image) {
	for j := car.posY; j < car.posY+100; j += 2 {
		carImgH.Move(fyne.NewPos(float32(j), float32(280)))
		time.Sleep(time.Millisecond)
	}
	carImgH.Move(fyne.NewPos(float32(2048), float32(280)))
}

func (p *Parking) reverseAbove(carImgV *canvas.Image, i int) {
	for j := p.posY[0]; j < p.posY[0]+100; j += 2 {
		carImgV.Move(fyne.NewPos(float32(p.posX[i]), float32(j)))
		time.Sleep(time.Millisecond)
	}
}

func (p *Parking) reverseBelow(carImgV *canvas.Image, i int) {
	for j := p.posY[1]; j > p.posY[1]-100; j -= 2 {
		carImgV.Move(fyne.NewPos(float32(p.posX[i]), float32(j)))
		time.Sleep(time.Millisecond)
	}
}
