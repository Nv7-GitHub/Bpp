package membuild

import (
	"github.com/Nv7-Github/Bpp/parser"
)

func IfBlock(p *Program, stm *parser.IfBlock) (Instruction, error) {
	cond, err := BuildStmt(p, stm.Condition)
	if err != nil {
		return nil, err
	}

	body := make([]Instruction, len(stm.Body))
	for i, stm := range stm.Body {
		stmt, err := BuildStmt(p, stm)
		if err != nil {
			return nil, err
		}
		body[i] = stmt
	}

	var el []Instruction
	if stm.Else != nil {
		el = make([]Instruction, len(stm.Else))
		for i, stm := range stm.Else {
			stmt, err := BuildStmt(p, stm)
			if err != nil {
				return nil, err
			}
			el[i] = stmt
		}
	}

	return func(p *Program) (Data, error) {
		cond, err := cond(p)
		if err != nil {
			return NewBlankData(), err
		}

		if cond.Value.(int) == 1 {
			for _, stmt := range body {
				out, err := stmt(p)
				if err != nil {
					return NewBlankData(), err
				}

				p.Runner(out)
			}
		} else if stm.Else != nil {
			for _, stmt := range el {
				out, err := stmt(p)
				if err != nil {
					return NewBlankData(), err
				}

				p.Runner(out)
			}
		}

		return NewBlankData(), nil
	}, nil
}
