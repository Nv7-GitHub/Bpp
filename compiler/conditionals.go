package compiler

import (
	"fmt"

	"github.com/Nv7-Github/Bpp/parser"
)

func CompileIf(val *parser.IfStmt) (string, parser.DataType, error) {
	cond, _, err := compileStmtRaw(val.Condition)
	if err != nil {
		return "", parser.NULL, err
	}
	body, dt, err := compileStmtRaw(val.Body)
	if err != nil {
		return "", parser.NULL, err
	}
	el, _, err := compileStmtRaw(val.Else)
	if err != nil {
		return "", parser.NULL, err
	}
	if dt == parser.NULL {
		body, _, err = compileStmt(val.Body)
		if err != nil {
			return "", parser.NULL, err
		}
		el, _, err = compileStmt(val.Else)
		if err != nil {
			return "", parser.NULL, err
		}
		return fmt.Sprintf("if ((%s) == 1) {%s} else {%s}", cond, body, el), parser.NULL, nil
	}
	if dt == parser.STRING {
		return fmt.Sprintf("((%s) == 1) ? (%s) : (%s)", cond, body, el), parser.STRING, nil
	}

	if dt == parser.FLOAT {
		return fmt.Sprintf("((%s) == 1) ? (float)(%s) : (float)(%s)", cond, body, el), parser.FLOAT, nil
	}

	return fmt.Sprintf("((%s) == 1) ? (int)(%s) : (int)(%s)", cond, body, el), parser.INT, nil
}

func CompileComparison(val *parser.ComparisonStmt) (string, parser.DataType, error) {
	left, ldt, err := compileStmtRaw(val.Left)
	if err != nil {
		return "", parser.NULL, err
	}
	right, rdt, err := compileStmtRaw(val.Right)
	if err != nil {
		return "", parser.NULL, err
	}
	if (ldt == parser.INT || rdt == parser.INT) && (ldt == parser.STRING || rdt == parser.STRING) { // Use atoi
		fmt.Println(ldt, rdt, parser.STRING)
		if ldt == parser.STRING {
			left = fmt.Sprintf("stoi(%s, &strsz)", left)
		}
		if rdt == parser.STRING {
			right = fmt.Sprintf("stoi(%s, &strsz)", right)
		}
	}
	return fmt.Sprintf("(bool)((%s) %s (%s))", left, compMap[val.Operation], right), parser.NULL, nil
}

var compMap = map[parser.Operator]string{
	parser.EQUAL:        "==",
	parser.NOTEQUAL:     "!=",
	parser.GREATER:      ">",
	parser.LESS:         "<",
	parser.GREATEREQUAL: ">=",
	parser.LESSEQUAL:    "<=",
}
