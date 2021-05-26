package compiler

import (
	"github.com/Nv7-Github/Bpp/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var m *ir.Module
var initBlock *ir.Block

func Compile(prog *parser.Program) (string, error) {
	m = ir.NewModule()
	tmpUsed = 0
	variables = make(map[string]Variable)
	autofree = make(map[value.Value]empty)
	generateBuiltins()

	main := m.NewFunc("main", types.I32, ir.NewParam("argc", types.I32), ir.NewParam("argv", types.NewPointer(types.I8Ptr)))
	initBlock = main.NewBlock("init")
	entry := main.NewBlock("entry")
	block := entry
	initMod(block)

	var err error
	block, err = CompileBlock(prog.Statements, block)
	if err != nil {
		return "", err
	}

	initBlock.NewBr(entry)

	for val := range autofree {
		block.NewCall(free, block.NewLoad(types.I8Ptr, val))
	}

	block.NewRet(constant.NewInt(types.I32, 0))

	return m.String(), nil
}

func CompileBlock(stms []parser.Statement, block *ir.Block) (*ir.Block, error) {
	var v value.Value
	var err error
	for _, val := range stms {
		v, block, err = CompileStmt(val, block)
		if err != nil {
			return block, err
		}

		if !val.Type().IsEqual(parser.NULL) {
			addLine(block, v)
		}
	}
	return block, nil
}

func addLine(block *ir.Block, val value.Value) {
	printVal(block, val)
	block.NewCall(printf, getStrPtr(newLine, block))
}

func printVal(block *ir.Block, val value.Value) {
	kind := val.Type()

	// Is array?
	_, ok := kind.(*types.PointerType)
	if ok {
		arrType, ok := kind.(*types.PointerType).ElemType.(*types.ArrayType)
		if ok {
			block.NewCall(printf, getStrPtr(openBracket, block))
			comma := getStrPtr(comma, block)
			for i := 0; i < int(arrType.Len); i++ {
				// Print all the vals
				ptr := block.NewGetElementPtr(arrType, val, constant.NewInt(types.I64, 0), constant.NewInt(types.I64, int64(i)))
				ld := block.NewLoad(arrType.ElemType, ptr)
				printVal(block, ld)
				if i != int(arrType.Len)-1 {
					block.NewCall(printf, comma)
				}
			}
			block.NewCall(printf, getStrPtr(closeBracket, block))
			return
		}
	}

	switch {
	case kind.Equal(types.I8Ptr):
		block.NewCall(printf, getStrPtr(strFmt, block), getStrPtr(val, block))

	case kind.Equal(types.Double):
		block.NewCall(printf, getStrPtr(fltFmt, block), val)

	case kind.Equal(types.I64):
		block.NewCall(printf, getStrPtr(intFmt, block), val)
	}
}

func addMalloc(len value.Value, block *ir.Block) value.Value {
	val := initBlock.NewAlloca(types.I8Ptr)
	//initBlock.NewStore(initBlock.NewCall(malloc, constant.NewInt(types.I64, 0)), val)
	initBlock.NewStore(constant.NewNull(types.I8Ptr), val)

	block.NewStore(block.NewCall(malloc, len), val)
	autofree[val] = empty{}
	return block.NewLoad(types.I8Ptr, val)
}
