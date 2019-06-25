package project2

import (
	"github.com/sirupsen/logrus"
	"strings"
)

type Args struct {
	Flags string
	Command string
	FlagsMapping map[string]string
	DefaultMapping map[string]string
	CommandMapping map[string]string
}


func (args *Args)Init(flags string,command string){
	args.Flags=flags
	args.Command=command
	args.FlagsMapping=make(map[string]string)
	args.CommandMapping=make(map[string]string)
	args.DefaultMapping=make(map[string]string)
	args.DefaultMapping["bool"]="false"
	args.DefaultMapping["int"]="0"
	args.DefaultMapping["string"]="."
}

func (args *Args)ParseFlags(){
	if !args.CheckInit(){
		logrus.Println("ParseFlags false!")
		return
	}
	s:=strings.Split(args.Flags,",")
	for i:=0;i<len(s);i++{
		args.FlagsMapping[s[i][0:1]]=s[i][2:]
	}
}

func (args *Args)ParseCommand(){
	if !args.CheckInit(){
		logrus.Println("ParseCommand false!")
		return
	}
	s:=strings.Split(args.Command," ")
	logrus.Println(s)
	for i:=0;i<len(s);i++{
		t:=strings.TrimSpace(s[i])
		logrus.Println("t:",t)
		if(t==""||t[0]!='-'){
			continue
		}
		if(args.CheckCommand(t)){
			tmpType,_:=args.FlagsMapping[t[1:]]
			if(i==len(s)-1){
				args.CommandMapping[t[1:]]=args.DefaultMapping[tmpType]
			} else {
				switch tmpType {
				case "string":
					if (args.StringParse(t[1:], strings.TrimSpace(s[i+1]))) {
						i++
						logrus.Println("string分解成功")
					}
				case "int":
					if (args.IntParse(t[1:], strings.TrimSpace(s[i+1]))) {
						i++
						logrus.Println("int分解成功")
					}
				case "bool":
					if (args.BoolParse(t[1:], strings.TrimSpace(s[i+1]))) {
						i++
						logrus.Println("nool分解成功")
					}
				}
			}
		}
	}
}

func (args *Args)GetValue(key string)string{
	if !args.CheckInit(){
		logrus.Println("GetValue false!")
		return ""
	}
	return args.CommandMapping[key]
}
