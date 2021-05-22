package parser

import (
	"fmt"
	"reflect"
	"strings"
)

func Parse(code string) (*Program, error) {
	code = strings.ReplaceAll(code, "\n\n", "\n") // Get rid of blank lines
	code = strings.TrimSpace(code)                // Remove blank spaces
	lns := strings.Split(code, "\n")

	functionTypes = make(map[string]FunctionType)

	prog := &Program{}
	scopes := NewScopeStack()
	scopes.AddScope(NewScope(prog))

	for i, val := range lns {
		stmt, err := ParseStmt(val, i+1, scopes)
		if err != nil {
			return nil, err
		}
		if stmt != nil {
			scopes.AddStatement(stmt)
		}
	}

	pScope := scopes.GetScope()
	p, ok := pScope.Block.(*Program)
	if !ok {
		return nil, fmt.Errorf("unterminated block: %s", reflect.TypeOf(pScope.Block))
	}
	scopes.FinishScope("", make([]Statement, 0))
	return p, nil
}

func ParseStmt(line string, num int, scope ...*ScopeStack) (Statement, error) {
	if strings.ContainsRune(line, '#') {
		line = line[:strings.IndexRune(line, '#')]
	}
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return nil, nil
	}
	if line[0] == '[' && line[len(line)-1] == ']' {
		split := strings.SplitN(line[1:], " ", 2)
		funcName := split[0]
		if len(split) == 1 {
			funcName = funcName[:len(funcName)-1]
		}

		parser, hasParser := parsers[funcName]
		var block Block
		var bParser BlockParser
		var fnType FunctionType
		isBParser := 0

		if !hasParser {
			// Is it a custom function
			var exists bool
			fnType, exists = functionTypes[funcName]
			if !exists {
				// No parser, is it a block end?
				if len(scope) > 0 {
					s := scope[0].GetScope()
					if !s.HasKeyword(funcName) {
						// Not a block end, is it a block start?
						bparser, exists := blocks[funcName]
						if exists {
							isBParser = 1
							bParser = bparser
						} else {
							return nil, fmt.Errorf("line %d: No such function '%s'", num, funcName)
						}
					} else {
						block = s.Block
					}
				} else {
					return nil, fmt.Errorf("line %d: No such function '%s'", num, funcName)
				}
			} else {
				isBParser = 2
			}
		}

		args := make([]string, 0)
		openedBrackets := 0
		openQuotation := false
		argTxt := strings.TrimSpace(line[len(funcName)+1 : len(line)-1])
		arg := ""

		for i, char := range argTxt {
			arg += string(char)

			switch char {
			case '[':
				openedBrackets++
			case ']':
				openedBrackets--
			case '"':
				openQuotation = !openQuotation
			}

			if (char == ' ' || i == len(argTxt)-1) && openedBrackets == 0 && !openQuotation {
				args = append(args, arg)
				arg = ""
				continue
			}
		}

		// Type checking
		argDat, err := ParseArgs(args, num)
		if err != nil {
			return nil, err
		}

		if hasParser {
			err = MatchTypes(argDat, num, parser.Signature)
			if err != nil {
				return nil, err
			}

			return parser.Parse(argDat, num)
		} else {
			if isBParser == 2 {
				err = MatchTypes(argDat, num, fnType.Signature)
				if err != nil {
					return nil, err
				}

				return &FunctionCallStmt{
					BasicStatement: &BasicStatement{line: num},
					ReturnType:     fnType.ReturnType,
					Name:           funcName,
					Args:           argDat,
				}, nil
			} else if isBParser == 1 {
				err = MatchTypes(argDat, num, bParser.Signature)
				if err != nil {
					return nil, err
				}

				block, err = bParser.Parse(argDat, num)
				if err != nil {
					return nil, err
				}

				s := NewScope(block)
				scope[0].AddScope(s)

				return nil, nil
			} else {
				err = MatchTypes(argDat, num, block.EndSignature())
				if err != nil {
					return nil, err
				}

				scope[0].FinishScope(funcName, argDat)

				return nil, nil
			}
		}
	}
	return ParseData(line, num), nil
}
