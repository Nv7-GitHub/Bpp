package builder

import (
	"fmt"
	"reflect"

	"github.com/Nv7-Github/bpp/old/ir"
)

func (b *builder) addInstruction(instr ir.Instruction) error {
	switch i := instr.(type) {
	case *ir.Const:
		return b.addConst(i)

	case *ir.AllocStatic:
		b.addAllocStatic(i)
		return nil

	case *ir.SetMemory:
		b.addSetMemory(i)
		return nil

	case *ir.GetMemory:
		b.addGetMemory(i)
		return nil

	case *ir.Math:
		b.addMath(i)
		return nil

	case *ir.Print:
		b.addPrint(i)
		return nil

	case *ir.Concat:
		b.addConcat(i)
		return nil

	case *ir.Array:
		b.addArray(i)
		return nil

	case *ir.AllocDynamic:
		b.addAllocDynamic(i)
		return nil

	case *ir.SetMemoryDynamic:
		b.addSetMemoryDynamic(i)
		return nil

	case *ir.GetMemoryDynamic:
		b.addGetMemoryDynamic(i)
		return nil

	case *ir.ArrayLength:
		b.addArrayLength(i)
		return nil

	case *ir.StringLength:
		b.addStringLength(i)
		return nil

	case *ir.Cast:
		b.addCast(i)
		return nil

	case *ir.Compare:
		b.addCompare(i)
		return nil

	case *ir.Jmp:
		b.addJmp(i)
		return nil

	case *ir.CondJmp:
		b.addCondJmp(i)
		return nil

	case *ir.JmpPoint:
		b.addJmpPoint()
		return nil

	case *ir.PHI:
		b.addPHI(i)
		return nil

	case *ir.StringIndex:
		b.addStringIndex(i)
		return nil

	case *ir.ArrayIndex:
		b.addArrayIndex(i)
		return nil

	case *ir.GetParam:
		return b.addGetParam(i)

	case *ir.FunctionCall:
		return b.addFunctionCall(i)

	case *ir.RandInt:
		b.addRandInt(i)
		return nil

	case *ir.RandFloat:
		b.addRandFloat(i)
		return nil

	case *ir.GetArg:
		b.addGetArg(i)
		return nil

	default:
		return fmt.Errorf("unknown instruction type: %s", reflect.TypeOf(i).String())
	}
}
