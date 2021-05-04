package parser

import (
	"fmt"
	"strings"
)

// Program is the main program, containing the source AST
type Program struct {
	Statements []Statement
}

func (p *Program) Type() DataType {
	return NULL
}

func (p *Program) Line() int {
	return 1
}

func Parse(code string) (Statement, error) {
	code = strings.ReplaceAll(code, "\n\n", "\n") // Get rid of blank lines
	lns := strings.Split(code, "\n")
	out := make([]Statement, len(lns))

	var err error
	for i, val := range lns {
		out[i], err = ParseStmt(val, i+1, i+1)
		if err != nil {
			return nil, err
		}
	}
	return &Program{
		Statements: out,
	}, nil
}

func ParseStmt(line string, num int, topLevel ...int) (Statement, error) {
	line = strings.TrimSpace(line)
	if line[0] == '[' && line[len(line)-1] == ']' {
		funcName := strings.SplitN(line[1:], " ", 2)[0]
		parser, exists := parsers[funcName]
		if !exists {
			return nil, fmt.Errorf("line %d: No such function '%s'", num, funcName)
		}

		args := make([]string, 0)
		openedBrackets := 0
		openQuotation := false
		argTxt := strings.TrimSpace(line[len(funcName)+1 : len(line)-1])
		arg := ""

		for i, char := range argTxt {
			if (char == ' ' || i == len(argTxt)-1) && openedBrackets == 0 && !openQuotation {
				if i == len(argTxt)-1 {
					arg += string(argTxt[len(argTxt)-1])
				}
				args = append(args, arg)
				arg = ""
				continue
			}
			arg += string(char)

			switch char {
			case '[':
				openedBrackets++
			case ']':
				openedBrackets--
			case '"':
				if openQuotation {
					openQuotation = true
				} else {
					openQuotation = false
				}
			}
		}

		// Type checking
		argDat, err := ParseArgs(args, num)
		if err != nil {
			return nil, err
		}
		err = MatchTypes(argDat, num)
		if err != nil {
			return nil, err
		}

		return parser.Parse(argDat, num)
	}
	return ParseData(line, num), nil
}
