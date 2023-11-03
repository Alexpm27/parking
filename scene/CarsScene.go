package scene

import (
	"estacionamiento/models"
	"fyne.io/fyne/v2"
	"math/rand"
	"time"
)

const (
	maxCapacity = 20
)

type CarsScene struct {
	parking   *models.Parking
	app       fyne.App
	win       fyne.Window
	container *fyne.Container
}

func newCarsScene(app fyne.App, win fyne.Window, container *fyne.Container) *CarsScene {
	return &CarsScene{
		parking:   models.NewParking(maxCapacity),
		app:       app,
		win:       win,
		container: container,
	}
}

func (m *CarsScene) CarGenerator() {
	go func() {
		for i := 0; i < 100; i++ {
			go m.addCar()
			time.Sleep(time.Duration(rand.ExpFloat64()/0.5) * time.Second)
		}
	}()
}

func (m *CarsScene) addCar() {
	car := models.NewCar()
	car.SetPosX(0)
	car.SetPosY(270)
	car.ImgV().Resize(fyne.NewSize(48, 96))
	car.ImgH().Resize(fyne.NewSize(96, 48))

	car.ImgH().Move(fyne.NewPos(0, 280))
	car.ImgV().Move(fyne.NewPos(0, 1003))

	m.container.Add(car.ImgV())
	m.container.Add(car.ImgH())
	m.container.Refresh()

	m.parking.ManageSpace(car, car.ImgV(), car.ImgH(), m.container)
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

	m.parking.Exit(car, car.ImgV(), car.ImgH(), m.container)
}
