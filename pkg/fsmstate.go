package txtfsm

type FSMJump struct {
	Label *FSMLabel
	Func  *FSMFunc
}

type FSMState struct {
	State        *FSMJump
	ErrHandler   *FSMJump
	ReturnStates []*FSMJump //Can be label or statefunc
}
