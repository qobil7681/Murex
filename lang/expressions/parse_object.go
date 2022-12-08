package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func (tree *expTreeT) createObjectAst(exec bool) error {
	// create JSON dict
	dt, nEscapes, err := tree.parseObject(exec)
	if err != nil {
		return err
	}
	tree.charPos -= nEscapes
	tree.appendAstWithPrimitive(symbols.ObjectBegin, dt)
	tree.charPos += nEscapes + 1
	return nil
}

func (tree *expTreeT) parseObject(exec bool) (*primitives.DataType, int, error) {
	var (
		nEscapes  int
		keyValueR [2][]rune
		keyValueI [2]interface{}
		stage     int
		obj       = make(map[string]interface{})
		//start     = tree.charPos
	)

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '\'', '"':
			// quoted string
			str, i, err := tree.parseString(r, exec)
			if err != nil {
				return nil, 0, err
			}
			keyValueR[stage&1] = append(keyValueR[stage&1], str...)
			nEscapes += i
			tree.charPos++

		case '%':
			switch tree.nextChar() {
			case '[', '{':
				// do nothing because action covered in the next iteration
			default:
				// string
				keyValueR[stage&1] = append(keyValueR[stage&1], r)
			}

		case '[':
			// start nested array
			if stage&1 == 0 {
				return nil, 0, fmt.Errorf("object keys cannot be an array")
			}
			dt, i, err := tree.parseArray(exec)
			if err != nil {
				return nil, 0, err
			}
			nEscapes += i
			keyValueI[1] = dt.Value
			tree.charPos++

		case '{':
			// start nested object
			if stage&1 == 0 {
				return nil, 0, raiseError(
					tree.expression, nil, tree.charPos, "object keys cannot be another object")
			}
			dt, i, err := tree.parseObject(exec)
			if err != nil {
				return nil, 0, err
			}
			nEscapes += i
			keyValueI[1] = dt.Value
			tree.charPos++

		case '$':
			switch {
			case isBareChar(tree.nextChar()):
				// inline scalar
				strOrVal := stage&1 == 0
				scalar, v, _, err := tree.parseVarScalar(exec, strOrVal)
				if err != nil {
					return nil, 0, err
				}
				if exec {
					keyValueI[stage&1] = v
				} else {
					keyValueI[stage&1] = string(scalar)
				}
			case tree.nextChar() == '{':
				// inline subshell
				strOrVal := stage&1 == 0
				subshell, v, _, err := tree.parseSubShell(exec, strOrVal)
				if err != nil {
					return nil, 0, err
				}
				if exec {
					keyValueI[stage&1] = v
				} else {
					keyValueI[stage&1] = string(subshell)
				}
			default:
				keyValueR[stage&1] = append(keyValueR[stage&1], r)
			}

		case '~':
			// tilde
			keyValueI[stage&1] = tree.parseVarTilde(true)

		case '@':
			// inline array
			if stage&1 == 0 {
				return nil, 0, raiseError(
					tree.expression, nil, tree.charPos, "arrays cannot be object keys")
			}
			name, _, err := tree.parseVarArray(exec)
			if err != nil {
				return nil, 0, err
			}
			return nil, 0, fmt.Errorf(
				"cannot expand an array into an object\nVariable name: @%s",
				string(name))

		case ':':
			if stage&1 == 1 {
				return nil, 0, raiseError(
					tree.expression, nil, tree.charPos, "invalid symbol ':' expecting ',' or '}' instead")
			}
			stage++
			if keyValueI[0] != nil {
				continue
			}
			keyValueI[0] = string(keyValueR[0])

		case '}', ',', '\n':
			if stage&1 == 0 {
				if r == '\n' {
					return nil, 0, raiseError(
						tree.expression, nil, tree.charPos,
						"unexpected new line, expecting ':' instead")
				}
				return nil, 0, raiseError(
					tree.expression, nil, tree.charPos, fmt.Sprintf(
						"invalid symbol '%s', expecting ':' instead",
						string(r)))
			}
			if keyValueI[0] == nil {
				return nil, 0, fmt.Errorf("object key cannot be null before %d", tree.charPos)
			}
			if len(keyValueR[1]) != 0 {
				keyValueI[1] = string(keyValueR[1])
			}

			s, err := types.ConvertGoType(keyValueI[0], types.String)
			if err != nil {
				return nil, 0, err
			}
			obj[s.(string)] = keyValueI[1]
			keyValueR = [2][]rune{nil, nil}
			keyValueI = [2]interface{}{nil, nil}
			stage++

			if r == '}' {
				goto endObject
			}

		case ' ', '\t', '\r':
			continue

		default:
			switch {
			case r == '-':
				next := tree.nextChar()
				if next < '0' || '9' < next {
					keyValueR[stage&1] = append(keyValueR[stage&1], r)
					continue
				}
				fallthrough
			case r >= '0' && '9' >= r:
				// number
				value := tree.parseNumber(r)
				tree.charPos--
				v, err := types.ConvertGoType(value, types.Number)
				if err != nil {
					return nil, 0, err
				}
				keyValueI[stage&1] = v

			default:
				// string
				keyValueR[stage&1] = append(keyValueR[stage&1], r)
			}
		}
	}

	return nil, 0, raiseError(
		tree.expression, nil, tree.charPos, "missing closing bracket (})")

endObject:
	tree.charPos--
	dt := &primitives.DataType{
		Primitive: primitives.Object,
		Value:     obj,
	}
	return dt, nEscapes, nil
}
