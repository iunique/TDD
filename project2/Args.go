package project2

import "github.com/sirupsen/logrus"

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
}

func (args *Args)ParseCommand(){
	if !args.Check(){
		logrus.Println("ParseCommand false!")
		return
	}
}

func (args *Args)GetValue(key byte)string{
	if !args.Check(){
		logrus.Println("GetValue false!")
		return ""
	}
	return args.CommandMapping[key]
}