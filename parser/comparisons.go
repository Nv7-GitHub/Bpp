package parser

import "fmt"

type Operator int

const (
	EQUAL        Operator = iota // =
	NOTEQUAL                     // !=
	GREATER                      // >
	LESS                         // <
	GREATEREQUAL                 // >=
	LESSEQUAL                    // <=
)

var opMap = map[string]Operator{
	"=":  EQUAL,
	"!=": NOTEQUAL,
	">":  GREATER,
	"<":  LESS,
	">=": GREATEREQUAL,
	"<=": LESSEQUAL,
}

// IfStmt is the equivalent of [IF stmt.Condition stmt.Body stmt.Else]
type IfStmt struct {
	*BasicStatement
	Condition Statement
	Body      Statement
	Else      Statement
}

func (i *IfStmt) Type() DataType {
	return ANY | NULL
}

// ComparisonStmt is the equivalent of [COMPARE stmt.Left stmt.Operation stmt.Right]
type ComparisonStmt struct {
	*BasicStatement
	Operation Operator
	Left      Statement
	Right     Statement
}

func (c *ComparisonStmt) Type() DataType {
	return INT
}

func SetupComparisons() {
	parsers["IF"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			return &IfStmt{
				Condition:      args[0],
				Body:           args[1],
				Else:           args[2],
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{INT, ANY | NULL, ANY | NULL},
	}

	parsers["COMPARE"] = StatementParser{
		Parse: func(args []Statement, line int) (Statement, error) {
			dat, ok := args[1].(*Data)
			if !ok {
				return nil, fmt.Errorf("line %d: argument 2 to COMPARE must be an operator", line)
			}
			opTxt, ok := dat.Data.(string)
			if !ok {
				return nil, fmt.Errorf("line %d: argument 2 to COMPARE must be an operator", line)
			}
			op, exists := opMap[opTxt]
			if !exists {
				return nil, fmt.Errorf("line %d: unknown comparison operator '%s'", line, opTxt)
			}
			return &ComparisonStmt{
				Operation:      op,
				Left:           args[0],
				Right:          args[2],
				BasicStatement: &BasicStatement{line: line},
			}, nil
		},
		Signature: []DataType{ANY | NULL, IDENTIFIER, ANY | NULL},
	}
}
