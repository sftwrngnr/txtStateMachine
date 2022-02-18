package txtfsm

type FSMState struct {
	Label        *FSMLabel
	StateFunc    *Funcs
	ErrHandler   *Funcs
	ReturnStates []interface{} //Can be label or statefunc
}
