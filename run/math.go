package run

import (
	"math"

	"github.com/Nv7-Github/bpp/ir"
	"github.com/Nv7-Github/bpp/parser"
)

func (r *Runnable) runMath(i *ir.Math) {
	typ := r.ir.Instructions[i.Val1].Type()
	if typ == ir.INT {
		switch i.Op {
		case parser.ADDITION:
			r.registers[r.Index] = r.registers[i.Val1].(int) + r.registers[i.Val2].(int)

		case parser.SUBTRACTION:
			r.registers[r.Index] = r.registers[i.Val1].(int) - r.registers[i.Val2].(int)

		case parser.MULTIPLICATION:
			r.registers[r.Index] = r.registers[i.Val1].(int) * r.registers[i.Val2].(int)

		case parser.DIVISION:
			r.registers[r.Index] = r.registers[i.Val1].(int) / r.registers[i.Val2].(int)

		case parser.POWER:
			r.registers[r.Index] = r.registers[i.Val1].(int) ^ r.registers[i.Val2].(int)
		}

		return
	}

	switch i.Op {
	case parser.ADDITION:
		r.registers[r.Index] = r.registers[i.Val1].(float64) + r.registers[i.Val2].(float64)

	case parser.SUBTRACTION:
		r.registers[r.Index] = r.registers[i.Val1].(float64) - r.registers[i.Val2].(float64)

	case parser.MULTIPLICATION:
		r.registers[r.Index] = r.registers[i.Val1].(float64) * r.registers[i.Val2].(float64)

	case parser.DIVISION:
		r.registers[r.Index] = r.registers[i.Val1].(float64) / r.registers[i.Val2].(float64)

	case parser.POWER:
		r.registers[r.Index] = math.Pow(r.registers[i.Val1].(float64), r.registers[i.Val2].(float64))
	}
}
