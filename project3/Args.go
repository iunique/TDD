package project3

import (
	"strconv"
	"strings"
)

type Args struct {
	Flags string
	Command string
	FlagsMapping map[ string ] string
	DefaultMapping map[ string ] string
	CommandMapping map[ string ] string
}
func (args *Args)Init(flags string, command string){
	args.Flags = flags
	args.Command = command
	args.FlagsMapping = make(map[ string ] string)
	args.CommandMapping = make(map[ string ] string)
	args.DefaultMapping = make(map [ string ] string)
	args.DefaultMapping[ "bool" ] = "false"
	args.DefaultMapping[ "int" ] = "0"
	args.DefaultMapping[ "string" ] = "default"
	args.FlagsMapping["l"] = "bool"
	args.FlagsMapping["s"] = "string"
	args.FlagsMapping["g"] = "int"

}
func (args *Args)ParseCommand(){
	s := strings.Split(args.Command, " ")//根据空格拆分
	for i := 0; i < len(s); i++{

		t := strings.TrimSpace( s[i] )//去掉首尾空格

		if( args.IsCommand( t ) ){//判断该词是否为命令
			tmpType, _ := args.FlagsMapping[ t[1:] ] //提取该命令对应的类型
			isSucc := args.TypeParse(tmpType, t[1:], strings.TrimSpace( s[i+1] ))//用TyParse去处理该词和下一个词是否匹配
			if isSucc { i++ }else{//若匹配，跳过下一个词
				args.CommandMapping[tmpType] = args.DefaultMapping[tmpType]
			}
		}
	}
}

func (args *Args)GetValue(key string) interface{} {

	switch args.FlagsMapping[ key ]{
		case "bool" :  result, _ := strconv.ParseBool( args.CommandMapping[ key ] ) ; return result
		case "int" : result, _ := strconv.Atoi(args.CommandMapping[ key ]) ; return result
		case "string" : return args.CommandMapping[ key ]
		default: return nil
	}
}