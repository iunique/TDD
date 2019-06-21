package project2

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type Args struct {
	Flags string
	Command string
	FlagsMapping map[byte]string
	CommandMapping map[byte]string
}

func (args *Args)Check()bool{
	if args.Flags==""||args.Command==""||args.FlagsMapping==nil||args.CommandMapping==nil{
		logrus.Println("the args is not ready")
		return false
	}
	return true
}

func (args *Args)Init(flags string,command string){
	args.Flags=flags
	args.Command=command
	args.FlagsMapping=make(map[byte]string)
	args.CommandMapping=make(map[byte]string)
}

func (args *Args)ParseFlags(){
	if !args.Check(){
		logrus.Println("ParseFlags false!")
		return
	}
	s:=strings.Split(args.Flags,",")
	for i:=0;i<len(s);i++{
		args.FlagsMapping[s[i][0]]=s[i][2:]
		//logrus.Printf("%c %s",s[i][0],s[i][2:])
	}
}

func (args *Args)ParseCommand(){
	if !args.Check(){
		logrus.Println("ParseCommand false!")
		return
	}
	s:=strings.Split(args.Flags,"-")
	for i:=0;i<len(s);i++{
		t:=strings.Split(s[i]," ")
		if(len(t)<2){
			if(args.FlagsMapping[t[0][0]]!=""){}
		}
	}
}

func (args *Args)CheckType(key byte,v interface{})bool{
	if !args.Check(){
		logrus.Println("GetValue false!")
		return false
	}
	value,isFind:=args.FlagsMapping[key]
	if(isFind){
		switch value {
			case "string":return "string"==reflect.TypeOf(v).String()
			case "int":return "int"==reflect.TypeOf(v).String()
			case "bool":return "bool"==reflect.TypeOf(v).String()
			default:logrus.Println("未匹配值")
		}
	}
	return false
}

func (args *Args)GetValue(key byte)string{
	if !args.Check(){
		logrus.Println("GetValue false!")
		return ""
	}
	return args.CommandMapping[key]
}