package parser

import "strconv"

func parseVariable(text string) Variable {
	if text[0] == '"' && text[len(text)-1] == '"' {
		return Variable{
			Type: STRING,
			Data: text[1 : len(text)-2],
		}
	}

	num, err := strconv.Atoi(text)
	if err == nil {
		return Variable{
			Type: INT,
			Data: num,
		}
	}

	flt, err := strconv.ParseFloat(text, 64)
	if err == nil {
		return Variable{
			Type: FLOAT,
			Data: flt,
		}
	}

	return Variable{
		Type: STRING | IDENTIFIER,
		Data: text,
	}
}
