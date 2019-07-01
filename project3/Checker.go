package project3


func (args *Args)IsCommand(s string) bool{
	if( len(s) == 2 && s[0] == '-' && s[1] >= 'a' && s[1] <= 'z' ){
		_, isExist := args.FlagsMapping[ s[1:] ]
		return isExist
	}
	return false
}
