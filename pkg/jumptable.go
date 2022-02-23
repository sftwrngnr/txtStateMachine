package txtfsm

const MaxStackSize = 20 //Prevent infinite recursion

type JumpElement struct {
	Current *FSMJump
	Next    *FSMJump
}

type JumpTable struct {
	StateLabel    string
	StateSequence []*JumpElement
}

type StateStack struct {
	Element []*JumpElement
}
