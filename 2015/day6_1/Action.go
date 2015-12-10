package main

type Action struct {
	Fn     CommandFn
	StartX int
	StartY int
	EndX   int
	EndY   int
}

func (action *Action) Execute(grid *[1000][1000]int) {
	action.Fn(grid, action.StartX, action.StartY, action.EndX, action.EndY)
}
