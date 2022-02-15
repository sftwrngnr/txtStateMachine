package txtfsm

type FSMState struct {
	StateFunc    *Funcs
	ErrHandler   *Funcs
	ReturnStates []*Funcs
}

type JumpElement struct {
	Current *Funcs
	Next    *Funcs
}

type JumpTable struct {
	InitState *FSMState
	CurState  *FSMState
	LabelMap  map[*Labels]*FSMState
}

type StateStack struct {
	Element []interface{}
}
