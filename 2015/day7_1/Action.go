package main

import (
	"log"
	"strings"
)

type Action struct {
	GateType string
	Operand1 int
	Operand2 int
	WriteTo  string
}

func NewAction() *Action {
	return &Action{}
}

func (action *Action) Execute(wires VariableContainer) {
	var result int

	switch action.GateType {
	case GATE_AND:
		result = action.GateAND()

	case GATE_LSHIFT:
		result = int(action.GATELShift())

	case GATE_NOT:
		result = action.GateNOT()

	case GATE_OR:
		result = action.GateOR()

	case GATE_RSHIFT:
		result = int(action.GATERShift())

	default:
		result = action.Operand2
	}

	log.Printf("Writing %d to %s\n", result, action.WriteTo)
	wires[strings.TrimSpace(action.WriteTo)] = int(result)
}

func (action *Action) GateAND() int {
	return action.Operand1 & action.Operand2
}

func (action *Action) GateOR() int {
	return action.Operand1 | action.Operand2
}

func (action *Action) GateNOT() int {
	return int(^uint16(action.Operand2))
}

func (action *Action) GATELShift() uint32 {
	return uint32(action.Operand1) << uint32(action.Operand2)
}

func (action *Action) GATERShift() uint32 {
	return uint32(action.Operand1) >> uint32(action.Operand2)
}
