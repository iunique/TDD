package project2

func (args *Args)StringParse(key string,v string)bool{
	defer func() {
		recover()
	}()
	if( args.CheckCommand( v ) ){
		args.CommandMapping[ key ] = args.DefaultMapping[ key ]
	} else {
		args.CommandMapping[ key ] = v
		return true
	}
	return false
}
