package txtfsm

type FSMInterface interface {
	pre_run(params []interface{}) (bool, error)
	run(params []interface{}) (int, error)
	post_run(params []interface{}) (bool, error)
}
