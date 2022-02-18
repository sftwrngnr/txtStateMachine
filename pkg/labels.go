package txtfsm

type ReturnAction struct {
	// Defines action for when last element of label is executed
}

type FSMLabel struct {
	Name     string
	Elements []*FSMElements
	Terminus *ReturnAction
}
