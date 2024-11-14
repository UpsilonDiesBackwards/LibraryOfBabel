package game

import "github.com/UpsilonDiesBackwards/LibraryOfBabel/game/player"

type App struct {
	Player *player.Player
}

func NewApplication() *App {
	return &App{
		Player: player.NewPlayer(),
	}
}

func (a *App) Initialise() {
	a.Player.Initialise()
}

func (a *App) Run() {
}

func (a *App) Destroy() {
}
