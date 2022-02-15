package txtfsm

type FSMPtype int

const (
	FSMVariable FSMPtype = iota
	FSMConstant
)

type FuncParam struct {
	ParamType FSMPtype
}
