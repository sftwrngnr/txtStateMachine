package txtfsm

type FSMState struct {
	Label        *FSMLabel
	StateFunc    *FSMFunc
	ErrHandler   *FSMFunc
	ReturnStates []interface{} //Can be label or statefunc
}
