package main

type StringStatus int

func (status StringStatus) String() string {
	if status == NAUGHTY {
		return "Naughty"
	}

	if status == NICE {
		return "Nice"
	}

	return "Unsure"
}
