package parser

import (
	"fmt"

	"github.com/Nv7-Github/bpp/types"
)

func Parse(code string, filename string) (*Program, error) {
	return ParseComplex(map[string]string{filename: code}, filename, make(map[string]ExternalFunction))
}

func ParseMultifile(files map[string]string, entryFile string) (*Program, error) {
	return ParseComplex(files, entryFile, make(map[string]ExternalFunction))
}

func ParseComplex(files map[string]string, entryFile string, externalFuncs map[string]ExternalFunction) (*Program, error) {
	prog := &Program{
		Functions:         make(map[string]*Function),
		VarTypes:          make(map[string]types.Type),
		ExternalFunctions: externalFuncs,
		Files:             files,
		Added:             map[string]empty{entryFile: {}},
	}
	code, exists := files[entryFile]
	if !exists {
		return nil, fmt.Errorf("entry file \"%s\" doesn't exist", entryFile)
	}
	built, err := prog.ParseCode(code, types.NewPos(entryFile))
	if err != nil {
		return nil, err
	}
	prog.Statements = built
	prog.Close()
	return prog, nil
}

func addMultifileParser() {
	parsers["IMPORT"] = Parser{
		Params: []types.Type{types.STRING},
		Parse: func(params []Statement, prog *Program, pos *types.Pos) (Statement, error) {
			filenameConst, ok := params[0].(*Const)
			if !ok {
				return nil, pos.NewError("filename must be constant")
			}
			file := filenameConst.Val.(string)

			_, exists := prog.Added[file]
			if exists {
				return nil, pos.NewError("circular imports not allowed")
			}

			code, exists := prog.Files[file]
			if !exists {
				return nil, pos.NewError("no such file \"%s\"", file)
			}

			built, err := prog.ParseCode(code, types.NewPos(file)) // This will add all functions and variables to the program on a global level
			if err != nil {
				return nil, err
			}

			// all code that is not functions or variables will be executed during the IMPORT
			prog.Added[file] = empty{}
			return &BlockStmt{
				BasicStmt: NewBasicStmt(pos),

				Body: built,
			}, nil
		},
	}
}
