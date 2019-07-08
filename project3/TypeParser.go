package project3

func (args *Args) TypeParse(tmpType string, key string, v string) bool {
	switch tmpType {
	case "string":
		return args.StringParse(key, v)
	case "int":
		return args.IntParse()
	case "bool":
		return args.BoolParse()
	}
	return false
}
