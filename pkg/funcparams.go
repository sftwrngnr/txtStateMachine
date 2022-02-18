package txtfsm

type FSMPtype int

const (
	FSMVariable FSMPtype = iota
	FSMConstant
)

type FuncParam struct {
	ParamType FSMPtype
	ParamData *Var //Constants are just specialized variables that are immutable when created
}

type FuncParams []*FuncParam
