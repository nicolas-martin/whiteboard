package f

type Delegate func(x int) int

type App struct {
	delegate Delegate
}

func (a *App) foo() {
	// ...
	y := a.delegate(1)
	_ = y
	// ...
}

func main() {
	delegate := func(x int) int {
		return x + 1
	}
	app := &App{delegate: delegate}
	app.foo()
}
