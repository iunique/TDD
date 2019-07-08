package project3

func (args *Args) StringParse(key string, v string) bool {
	result := args.IsCommand(v)
	if !result {
		args.CommandMapping[key] = v
	}
	return result
}
