package expressions

func (tree *expTreeT) Dump() interface{} {
	var (
		dump  = make(map[string]interface{})
		nodes = make([]interface{}, len(tree.ast))
	)

	for i := range tree.ast {
		node := make(map[string]interface{})
		node["key"] = tree.ast[i].key.String()
		node["value"] = tree.ast[i].Value()
		node["pos"] = tree.ast[i].pos
		if tree.ast[i].dt != nil {
			node["dt.prim"] = tree.ast[i].dt.Primitive.String()
			node["dt.murex"] = tree.ast[i].dt.DataType()
			node["dt.value"] = tree.ast[i].dt.Value
		} else {
			node["dt"] = "unset"
		}

		nodes[i] = node
	}

	dump["ast"] = nodes
	dump["charPos"] = tree.charPos
	dump["charOffset"] = tree.charOffset
	dump["astPos"] = tree.astPos
	dump["expression"] = string(tree.expression)

	return dump
}