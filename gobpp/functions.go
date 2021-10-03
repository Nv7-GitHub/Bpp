package gobpp

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

type empty struct{}

var hasReturn = make(map[string]empty)

var typeMap = map[string]string{
	"string":  "STRING",
	"float64": "FLOAT",
	"float32": "FLOAT",
	"int32":   "INT",
	"int64":   "INT",
	"int":     "INT",
}

var fnRetTypes = make(map[string]string)

func (p *Program) addFuncDecl(fn *ast.FuncDecl) error {
	name := strings.ToUpper(fn.Name.Name)
	p.FuncName = name

	if fn.Type.Results != nil {
		typeName := fn.Type.Results.List[0].Type.(*ast.Ident).Name
		kind, exists := typeMap[typeName]
		if !exists {
			return fmt.Errorf("%s: unknown type %s", p.NodePos(fn), typeName)
		}
		fnRetTypes[name] = kind
	}

	_, _ = p.WriteString("[FUNCTION ")
	_, _ = p.WriteString(name)

	for _, arg := range fn.Type.Params.List {
		var kind string

		switch v := arg.Type.(type) {
		case *ast.Ident:
			kind = typeMap[v.Name]

		case *ast.ArrayType:
			kind = "ARRAY"

		default:
			return fmt.Errorf("%s: unknown function parameter type: %s", p.NodePos(arg), reflect.TypeOf(arg.Type))
		}

		_, _ = fmt.Fprintf(p, " [PARAM %s %s]", arg.Names[0].Name, kind)
	}
	_, _ = p.WriteString("]\n")

	err := p.AddBlock(fn.Body.List)
	if err != nil {
		return err
	}

	_, exists := hasReturn[name]
	if !exists {
		_, _ = p.WriteString("[RETURN NULL [NULL]]\n")
	}
	p.FuncName = ""
	return nil
}

// AddBlock adds a block of statements to a B++ program, add a boolean parameter to the end of the call to remove the "\n" from the end of the code generated
func (p *Program) AddBlock(args []ast.Stmt) error {
	var err error
	for _, stmt := range args {
		err = p.AddStmt(stmt)
		if err != nil {
			return err
		}
		_, _ = p.WriteString("\n")
	}
	return nil
}

func (p *Program) addReturnStmt(s *ast.ReturnStmt) error {
	_, _ = p.WriteString("[RETURN ")
	kind, exists := fnRetTypes[p.FuncName]
	if !exists {
		return fmt.Errorf("%s: no return type", p.NodePos(s))
	}
	_, _ = p.WriteString(kind)
	_, _ = p.WriteString(" ")

	err := p.AddExpr(s.Results[0])
	if err != nil {
		return err
	}

	_, _ = p.WriteString("]")

	hasReturn[p.FuncName] = empty{}

	return nil
}
