package run

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/bpp/ir"
)

func (r *Runnable) runInstruction(index int) error {
	instr := r.ir.Instructions[index]
	switch i := instr.(type) {
	case *ir.Const:
		r.runConst(index)
		return nil

	case *ir.Print:
		return r.runPrint(i)

	case *ir.AllocStatic:
		r.runAllocStatic(index, i)
		return nil

	case *ir.AllocDynamic:
		r.runAllocDynamic(index, i)
		return nil

	case *ir.SetMemory:
		r.runSetMemory(i)
		return nil

	case *ir.SetMemoryDynamic:
		r.runSetMemoryDynamic(i)
		return nil

	case *ir.GetMemory:
		r.runGetMemory(index, i)
		return nil

	case *ir.GetMemoryDynamic:
		r.runGetMemoryDynamic(index, i)
		return nil

	case *ir.Math:
		r.runMath(index, i)
		return nil

	case *ir.Cast:
		return r.runCast(index, i)

	default:
		return fmt.Errorf("unknown instruction type: %s", reflect.TypeOf(i).String())
	}
}
