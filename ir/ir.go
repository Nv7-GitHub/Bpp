package ir

import (
	"fmt"

	"github.com/Nv7-Github/bpp/parser"
)

func NewIR() *IR {
	ir := &IR{
		Functions: make([]Function, 0),
		Fns:       make(map[string]int),
	}
	return ir
}

func CreateIR(prog *parser.Program) (*IR, error) {
	ir := NewIR()
	// Add functions
	for i, stm := range prog.Statements {
		f, ok := stm.(*parser.FunctionBlock)
		if !ok {
			continue
		}
		err := ir.AddFunction(f)
		if err != nil {
			return nil, err
		}

		// Remove from array
		copy(prog.Statements[i:], prog.Statements[i+1:])
		prog.Statements = prog.Statements[:len(prog.Statements)-1]
	}

	// Add statements
	ir.Instructions = make([]Instruction, 0)
	ir.vars = make(map[string]varData)
	for _, stmt := range prog.Statements {
		_, err := ir.AddStmtTop(stmt)
		if err != nil {
			return nil, err
		}
	}
	return ir, nil
}

type Print struct {
	Val int
}

func (p *Print) String() string {
	return fmt.Sprintf("Print: %d", p.Val)
}

func (p *Print) Type() Type {
	return NULL
}

func (i *IR) newPrint(val int) int {
	return i.AddInstruction(&Print{Val: val})
}

func (i *IR) AddStmtTop(stmt parser.Statement) (int, error) {
	ind, err := i.AddStmt(stmt)
	if err != nil {
		return 0, err
	}
	typ := i.GetInstruction(ind).Type()
	if typ != NULL {
		switch typ {
		case STRING:
			i.newPrint(ind)

		case INT, FLOAT:
			casted := i.newCast(ind, STRING)
			i.newPrint(casted)
		}
	}
	return ind, nil
}
