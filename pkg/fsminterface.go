package txtfsm

type FSMInterface interface {
	run(params []interface{}) (int, error)
}
