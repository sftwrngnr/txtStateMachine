package txtfsm

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getReservedList() []string {
	return []string{":EXEC", ":DONE", ":RETURN", ":CONTINUE", ":VARS", ":ABORT", ":RESTART"}
}

func getTerminusList() []string {
	return []string{":RETURN", ":CONTINUE", ":ABORT", ":RESTART", ":DONE"}
}

type FSMElements struct {
	FSMLabels []FSMLabel
	FSMFuncs  []Funcs
	FSMVars   []Var
}

type FSMParser struct {
	FsmText []string
	Pmap    FSMElements
}

func (f *FSMParser) ReadFile(fn string) error {
	data, err := os.ReadFile(fn)
	if err != nil {
		return err
	}
	f.FsmText = strings.Split(string(data), "\n")
	return nil
}

func (f *FSMParser) trimComment(s string) string {
	if strings.Contains(s, "#") {
		return s[0:strings.Index(s, "#")]
	}
	return s
}

func (f *FSMParser) checkReserved(s string) bool {
	for _, v := range getReservedList() {
		if strings.Contains(s, v) {
			fmt.Printf("Reserved word: %s\n", v)
			return true
		}
	}
	return false
}

func (f *FSMParser) checkTerminus(s string) bool {
	for _, v := range getTerminusList() {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}

func (f *FSMParser) getVarName(v, s string) string {
	retval := s[0:strings.Index(s, v)]
	return strings.Trim(retval, " \t")
}

func (f *FSMParser) getVarIntVal(s string) (int, error) {
	vIdx := strings.Index(s, "INT")
	if vIdx > -1 {
		vIdx += 3
		cval := strings.Trim(s[vIdx:], " \t\n")
		if len(cval) == 0 {
			return 0, nil
		}
		rv, err := strconv.Atoi(cval)
		if err != nil {
			return 0, err // Failed to parse
		}
		return rv, nil
	}
	return 0, nil
}

func (f *FSMParser) getVarFloatVal(s string) (float32, error) {
	vIdx := strings.Index(s, "FLOAT")
	if vIdx > -1 {
		vIdx += 5
		cval := strings.Trim(s[vIdx:], " \t\n")
		if len(cval) == 0 {
			return float32(0), nil
		}
		rv, err := strconv.ParseFloat(cval, 32)
		if err != nil {
			return 0, err // Failed to parse
		}
		return float32(rv), nil
	}
	return 0, nil
}

func (f *FSMParser) getVarStringVal(s string) (string, error) {
	vIdx := strings.Index(s, "STRING")
	if vIdx > -1 {
		vIdx += 6
		cval := strings.Trim(s[vIdx:], " \t\n")
		return cval, nil
	}
	return "", nil

}

func (f *FSMParser) processVarLine(s string) {
	// Insure that there is a valid variable type
	ValidTypes := []string{"INT", "FLOAT", "STRING"}
	for _, v := range ValidTypes {
		if strings.Contains(s, v) {
			// We know which one this is!
			var tvar Var
			tvar.Name = f.getVarName(v, s)
			switch v {
			case "INT":
				// Process int
				tvar.VType = VInt
				tv, err := f.getVarIntVal(s)
				if err != nil {
					log.Fatal(err)
				}
				tvar.curval = tv

			case "FLOAT":
				// Process float
				tvar.VType = VFloat
				tv, err := f.getVarFloatVal(s)
				if err != nil {
					log.Fatal(err)
				}
				tvar.curval = tv
			case "STRING":
				// Process string
				tvar.VType = VString
				tv, err := f.getVarStringVal(s)
				if err != nil {
					log.Fatal(err)
				}
				tvar.curval = tv

			}
			f.Pmap.FSMVars = append(f.Pmap.FSMVars, tvar)
		}
	}
}

func (f *FSMParser) processVars(i int) (int, error) {
	for n, v := range f.FsmText[i:] {
		if f.checkTerminus(v) {
			return n + i, nil
		}
		f.processVarLine(v)
	}
	return i, nil
}

func (f *FSMParser) BuildFSMElements() (bool, error) {
	// Iterate through state file, build variable list, label list, and function list
	var SkipLine int
	var err error
	for i, v := range f.FsmText {
		if len(v) == 0 {
			continue
		}
		if v[0] == '#' {
			continue
		}
		fmt.Printf("i %d, SkipLine %d\n", i, SkipLine)
		if i < SkipLine {
			continue
		}
		if strings.Contains(v, ":VARS") {
			fmt.Printf("Found Vars tag!\n")
			SkipLine, err = f.processVars(i + 1)
			fmt.Printf("Variables are: %v\n", f.Pmap.FSMVars)
			if err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

func (f *FSMParser) processLine(l int) {

}

func (f *FSMParser) Parse() (bool, error) {
	if succ, err := f.BuildFSMElements(); (err == nil) && succ {
		/*
			if len(f.FsmText) > 0 {
				newI := -1
				for l, v := range f.FsmText {
					if newI > l {
						continue
					}
					if len(v) == 0 {
						continue
					}
					if v[0] == '#' {
						continue
					}
					//newL := f.trimComment(v)
					//fmt.Printf("%d %s\n", l, newL)
					f.processLine(l)
				}
			}
		*/

	} else {
		return succ, err
	}
	return true, nil
}
