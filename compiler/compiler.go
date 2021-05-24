package compiler

import (
	"github.com/Nv7-Github/Bpp/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

var m *ir.Module

func Compile(prog *parser.Program) (string, error) {
	m = ir.NewModule()
	tmpUsed = 0
	generateBuiltins()

	main := m.NewFunc("main", types.I32)
	block := main.NewBlock("entry")

	var err error
	block, err = CompileBlock(prog.Statements, block)
	if err != nil {
		return "", err
	}

	block.NewRet(constant.NewInt(types.I32, 0))

	return m.String(), nil
}

type Builder func(parser.Statement, *ir.Block) (value.Value, *ir.Block, error)

var builders = make(map[string]Builder)

func CompileBlock(stms []parser.Statement, block *ir.Block) (*ir.Block, error) {
	var v value.Value
	var err error
	for _, val := range stms {
		v, block, err = CompileStmt(val, block)
		if err != nil {
			return block, err
		}

		addLine(block, v, val.Type())
	}
	return block, nil
}
