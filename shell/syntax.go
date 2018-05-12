package shell

func syntaxCompletion(line []rune, pos int) ([]rune, int) {
	pt, _ := parse(line)
	switch {
	case pt.QuoteSingle && pt.QuoteBrace == 0 && pt.NestedBlock == 0:
		if pos < len(line)-1 && line[pos] == '\'' && line[len(line)-1] == '\'' {
			return line[:len(line)-1], pos
		}
		if pos < len(line)-1 || line[pos] != '\'' {
			return append(line, '\''), pos
		}

	case pt.QuoteDouble && pt.QuoteBrace == 0 && pt.NestedBlock == 0:
		if pos < len(line)-1 && line[pos] == '"' && line[len(line)-1] == '"' {
			return line[:len(line)-1], pos
		}
		if pos < len(line)-1 || line[pos] != '"' {
			return append(line, '"'), pos
		}

	case pt.QuoteBrace > 0 && pt.NestedBlock == 0:
		if pos < len(line)-1 || line[pos] != '(' {
			return append(line, ')'), pos
		}

	case pt.QuoteBrace < 0:
		if line[pos] == ')' && line[len(line)-1] == ')' && pos != len(line)-1 {
			return line[:len(line)-1], pos
		}

	case pt.NestedBlock > 0 && pt.QuoteBrace == 0:
		if pos < len(line)-1 || line[pos] != '{' {
			return append(line, '}'), pos
		}

	case pt.NestedBlock < 0:
		if line[pos] == '}' && line[len(line)-1] == '}' && pos != len(line)-1 {
			return line[:len(line)-1], pos
		}

	case pos > 0 && len(line) > pos && line[pos-1] == '[':
		if pos < len(line)-1 {
			r := append(line[:pos+1], ']')
			return append(r, line[pos+2:]...), pos
		}
		return append(line, ']'), pos

	}
	return line, pos
}
