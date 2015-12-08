package main

type Santa struct {
	GPS  *GPS
	Name string
}

func NewSanta(name string, housesVisited int) *Santa {
	return &Santa{
		GPS:  NewGPS(),
		Name: name,
	}
}

func (santa *Santa) Move(direction string) Point {
	if direction == MOVE_UP {
		santa.GPS.MoveUp()
	}

	if direction == MOVE_DOWN {
		santa.GPS.MoveDown()
	}

	if direction == MOVE_LEFT {
		santa.GPS.MoveLeft()
	}

	if direction == MOVE_RIGHT {
		santa.GPS.MoveRight()
	}

	return santa.GetPosition()
}

func (santa *Santa) GetPosition() Point {
	return santa.GPS.Location
}
