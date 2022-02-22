package txtfsm

type FSMPType int
type FSMDType int

const (
	FSMVariable FSMPType = iota
	FSMConstant
)

const (
	FSMInt FSMDType = iota
	FSMFloat
	FSMString
)

type FuncParam struct {
	ParamType     FSMPType
	ParamDataType FSMDType
}

type FuncParams []*FuncParam

func (fp FuncParams) Params(ftp FSMPType, fsmt FSMDType) FuncParams {
	return append(fp, &FuncParam{ParamType: ftp, ParamDataType: fsmt})
}
