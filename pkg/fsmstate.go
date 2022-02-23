package txtfsm

type FSMJump struct {
	Label *FSMLabel
	Func  *FSMFunc
}

type FSMState struct {
	State        *FSMJump
	Params       *FuncParams
	ErrHandler   *FSMJump
	ReturnStates []*FSMJump //Can be label or statefunc
	Next         *FSMJump
}
