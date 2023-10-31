package i

type Delegate interface {
	Do(x int) int
}

type App struct {
	delegate Delegate
}

func (a *App) foo() {
	// ...
	y := a.delegate.Do(1)
	_ = y
	// ...
}

type DelegateImpl struct{}

func (d *DelegateImpl) Do(x int) int {
	return x + 1
}

func main() {
	delegate := &DelegateImpl{}
	app := &App{delegate}
	app.foo()

}
