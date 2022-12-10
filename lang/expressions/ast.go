package expressions

import (
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

type astNodeT struct {
	key   symbols.Exp
	value []rune
	pos   int
	dt    *primitives.DataType
}

func (node *astNodeT) Value() string {
	return string(node.value)
}

type ParserT struct {
	ast          []*astNodeT
	charPos      int
	charOffset   int
	astPos       int
	expression   []rune
	isSubExp     bool
	p            *lang.Process
	strictArrays interface{}
	expandGlob   interface{}
	Statement    *StatementT
}

func (tree *ParserT) nextChar() rune {
	if tree.charPos+1 >= len(tree.expression) {
		return 0
	}
	return tree.expression[tree.charPos+1]
}

func (tree *ParserT) prevChar() rune {
	if tree.charPos < 1 {
		return 0
	}
	return tree.expression[tree.charPos-1]
}

func (tree *ParserT) appendAst(key symbols.Exp, value ...rune) {
	tree.ast = append(tree.ast, &astNodeT{
		key:   key,
		value: value,
		pos:   tree.charPos + tree.charOffset - len(value),
	})
}

func (tree *ParserT) appendAstWithPrimitive(key symbols.Exp, dt *primitives.DataType) {
	tree.ast = append(tree.ast, &astNodeT{
		key: key,
		dt:  dt,
		pos: tree.charPos + tree.charOffset,
	})
}

func (tree *ParserT) foldAst(new *astNodeT) error {
	switch {
	case tree.astPos <= 0:
		return fmt.Errorf("cannot fold when tree.astPos<%d> <= 0<%d> (%s)",
			tree.astPos, len(tree.ast), consts.IssueTrackerURL)

	case tree.astPos >= len(tree.ast)-1:
		return fmt.Errorf("cannot fold when tree.astPos<%d> >= len(tree.ast)-1<%d> (%s)",
			tree.astPos, len(tree.ast), consts.IssueTrackerURL)

	case len(tree.ast) == 3:
		tree.ast = []*astNodeT{new}

	case tree.astPos == 1:
		tree.ast = append([]*astNodeT{new}, tree.ast[3:]...)

	case tree.astPos == len(tree.ast)-2:
		tree.ast = append(tree.ast[:len(tree.ast)-3], new)

	default:
		start := append(tree.ast[:tree.astPos-1], new)
		end := tree.ast[tree.astPos+2:]
		tree.ast = append(start, end...)
	}

	return nil
}

// memory safe
func (tree *ParserT) prevSymbol() *astNodeT {
	if tree.astPos-1 < 0 {
		return nil
	}

	return tree.ast[tree.astPos-1]
}

// memory safe
func (tree *ParserT) currentSymbol() *astNodeT {
	if tree.astPos < 0 || tree.astPos >= len(tree.ast) {
		return nil
	}

	return tree.ast[tree.astPos]
}

// memory safe
func (tree *ParserT) nextSymbol() *astNodeT {
	if tree.astPos+1 >= len(tree.ast) {
		return nil
	}

	return tree.ast[tree.astPos+1]
}

func (tree *ParserT) getLeftAndRightSymbols() (*astNodeT, *astNodeT, error) {
	left := tree.prevSymbol()
	right := tree.nextSymbol()

	if left == nil {
		return nil, nil, raiseError(tree.expression, tree.ast[tree.astPos], 0, "missing value left of operation")
	}

	if right == nil {
		return nil, nil, raiseError(tree.expression, tree.ast[tree.astPos], 0, "missing value right of operation")
	}

	return left, right, nil
}

func node2primitive(node *astNodeT) (*primitives.DataType, error) {
	switch node.key {
	case symbols.Number:
		f, err := strconv.ParseFloat(node.Value(), 64)
		if err != nil {
			return nil, raiseError(nil, node, 0, err.Error())
		}
		return &primitives.DataType{
			Primitive: primitives.Number,
			Value:     f,
		}, nil

	case symbols.QuoteSingle:
		return &primitives.DataType{
			Primitive: primitives.String,
			Value:     node.Value(),
		}, nil

	case symbols.QuoteDouble:
		// TODO: expand vars
		return &primitives.DataType{
			Primitive: primitives.String,
			Value:     node.Value(),
		}, nil

	case symbols.Boolean:
		return &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     types.IsTrueString(string(node.value), 0),
		}, nil

	case symbols.Bareword:
		return &primitives.DataType{
			Primitive: primitives.Bareword,
			Value:     nil,
		}, nil

	case symbols.Calculated:
		return &primitives.DataType{
			Primitive: primitives.Null,
			Value:     nil,
		}, nil
	}

	return nil, raiseError(nil, node, 0, fmt.Sprintf("unexpected error converting node to primitive (%s)", consts.IssueTrackerURL))
}

func (tree *ParserT) StrictArrays() bool {
	if tree.strictArrays != nil {
		return tree.strictArrays.(bool)
	}

	var err error
	tree.strictArrays, err = tree.p.Config.Get("proc", "strict-arrays", "bool")
	if err != nil {
		tree.strictArrays = true
	}

	return tree.strictArrays.(bool)
}

func (tree *ParserT) ExpandGlob() bool {
	if tree.expandGlob != nil {
		return tree.expandGlob.(bool)
	}

	var err error
	tree.expandGlob, err = tree.p.Config.Get("shell", "expand-glob", "bool")
	if err != nil {
		tree.expandGlob = true
	}

	return tree.expandGlob.(bool)
}
