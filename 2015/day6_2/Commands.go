package main

const (
	COMMAND_TURN_ON  string = "turn on"
	COMMAND_TURN_OFF string = "turn off"
	COMMAND_TOGGLE   string = "toggle"
)

type CommandFn func(*[1000][1000]int, int, int, int, int)

func TurnOn(grid *[1000][1000]int, startX, startY, endX, endY int) {
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			grid[x][y] = grid[x][y] + 1
		}
	}
}

func TurnOff(grid *[1000][1000]int, startX, startY, endX, endY int) {
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			if grid[x][y] > 0 {
				grid[x][y] = grid[x][y] - 1
			}
		}
	}
}

func Toggle(grid *[1000][1000]int, startX, startY, endX, endY int) {
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			grid[x][y] = grid[x][y] + 2
		}
	}
}
