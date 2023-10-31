package main

import (
	"fmt"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			fx.Annotate(
				newLittleBoy,
				fx.ResultTags(`name:"littleBoy"`),
			),

			fx.Annotate(
				newLittleBoy2,
				fx.ResultTags(`name:"littleBoy3"`),
			),
			newLittleCat,
			fx.Annotate(
				newBig,
			),
			newOther,
		),
		fx.Invoke(DoBig),
	).Run()
}

type Other struct {
	Lc ILittleCat
}

func newOther(lc ILittleCat) *Other {
	return &Other{Lc: lc}
}

type Big struct {
	Lb  ILittleBoy
	Lb2 ILittleBoy
	Lc  ILittleCat
}
type Args struct {
	fx.In
	Ib  ILittleBoy `name:"littleBoy"`
	Ib2 ILittleBoy `name:"littleBoy3"`
	Lc  ILittleCat
}

func newBig(args Args) *Big {
	fmt.Println()
	return &Big{Lb: args.Ib, Lb2: args.Ib2, Lc: args.Lc}

}

func DoBig(b *Big, o *Other) {
	b.Lb.DoBoyThing()
	b.Lb2.DoBoyThing()
	b.Lc.DoCatThing()
	_ = o
}

type ILittleBoy interface {
	DoBoyThing()
}

type ILittleCat interface {
	DoCatThing()
}

type LittleBoy struct{}

func newLittleBoy() ILittleBoy {
	return &LittleBoy{}
}

func (lb *LittleBoy) DoBoyThing() { fmt.Println("little boy") }

type LittleBoy2 struct{}

func newLittleBoy2() ILittleBoy {
	return &LittleBoy2{}
}

func (lb2 *LittleBoy2) DoBoyThing() { fmt.Println("little boy2") }

type LittleCat struct{}

func newLittleCat() ILittleCat {
	return &LittleCat{}
}

func (lb *LittleCat) DoCatThing() { fmt.Println("little cat") }
