package scene

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type MainScene struct {
	app            fyne.App
	win            fyne.Window
	sceneContainer *fyne.Container
}

func NewMainScene() *MainScene {
	myApp := app.New()
	appWindow := myApp.NewWindow("Park")
	appWindow.Resize(fyne.NewSize(1366, 768))
	return &MainScene{
		app:            myApp,
		win:            appWindow,
		sceneContainer: container.NewWithoutLayout(),
	}
}

func (m *MainScene) Start() {
	background := canvas.NewImageFromFile("img/parking.png")
	background.Resize(fyne.NewSize(1200, 668))
	background.Move(fyne.NewPos(100, 50))
	m.sceneContainer.Add(background)
	m.win.SetContent(m.sceneContainer)
	carsScene := newCarsScene(m.app, m.win, m.sceneContainer)
	go carsScene.CarGenerator()

	m.win.CenterOnScreen()
	m.win.ShowAndRun()
}
