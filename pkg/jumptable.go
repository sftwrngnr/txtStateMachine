package txtfsm

type FSMElements struct {
	labels []Labels
	funcs  []Funcs
	vars   []Vars
}

type FSMState struct {
	StateFunc    *Funcs
	ErrHandler   *Funcs
	ReturnStates []*Funcs
}

type JumpTable struct {
	InitState *FSMState
	CurState  *FSMState
	LabelMap  map[*Labels]*FSMState
}

type StateStack struct {
	States []*FSMState
}
