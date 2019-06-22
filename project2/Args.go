package project2

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type Args struct {
	Flags string
	Command string
	FlagsMapping map[string]string
	CommandMapping map[string]string
}

func (args *Args)Check()bool{
	if args.Flags==""||args.Command==""||args.FlagsMapping==nil||args.CommandMapping==nil{
		logrus.Errorf("the args is not ready")
		return false
	}
	return true
}

func (args *Args)Init(flags string,command string){
	args.Flags=flags
	args.Command=command
	args.FlagsMapping=make(map[string]string)
	args.CommandMapping=make(map[string]string)
}

func (args *Args)ParseFlags(){
	if !args.Check(){
		logrus.Println("ParseFlags false!")
		return
	}
	s:=strings.Split(args.Flags,",")
	for i:=0;i<len(s);i++{
		args.FlagsMapping[s[i][0:1]]=s[i][2:]
	}
}

func (args *Args)ParseCommand(){
	if !args.Check(){
		logrus.Println("ParseCommand false!")
		return
	}
	s:=strings.Split(args.Command,"-")
	logrus.Println(s)
	for i:=0;i<len(s);i++{
		t:=strings.Split(s[i]," ")
		if(len(t)==1){//若拆分后没有参数，则进行默认初始化
			_,isExist:=args.FlagsMapping[t[0]]
			if(!isExist){
				if (IsDigit(t[0])) { //-d后面接负数情况
					args.CommandMapping["d"]="-"+t[0]
				} else {
					logrus.Println("commond 字符不存在！输入command错误")
					panic("commond 字符不存在！输入command错误")
				}
			}
			if(t[0]=="l"){
				args.CommandMapping["l"]="false"
			}else if(t[0]=="d"){
				args.CommandMapping["d"]="0"
			}else if(t[0]=="f"){
				args.CommandMapping["f"]="."
			}
		} else if(len(t)==2){//拆分后有参数且只能跟一个参数
			//检查第一个字符是否合法
			_,isExist:=args.FlagsMapping[t[0]]
			if(!isExist){
				logrus.Println("commond 字符不存在！输入command错误")
				panic("commond 字符不存在！输入command错误")
			}
			//检查输入的参数是否和值匹配
			if(t[0]=="l"){
				if(t[1]=="false"||t[1]=="true") { //默认小写
					args.CommandMapping[t[0]]=t[1]
				}else {
					logrus.Println("commond 字符不合法！输入command错误")
					panic("commond 字符不合法！输入command错误")
				}
			}else if(t[0]=="d"){
				if (IsDigit(t[1])) { //
					args.CommandMapping["d"]=t[1]
				} else {
					logrus.Println("commond 字符不合法！输入command错误")
					panic("commond 字符不合法！输入command错误")
				}
			}else if(t[0]=="f"){
				args.CommandMapping["f"]=t[1]
			}

		} else{
			logrus.Println("commond 字符不合法！输入command错误")
			panic("commond 字符不合法！输入command错误")
		}
	}
}
//判断个位数字
func IsSingleDigit(data string) bool {

	digit := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for _, item := range digit {
		if data == item {
			return true
		}
	}
	return false
}
//判断数字
func IsDigit(data string) bool{
	for _, item := range data{
		if IsSingleDigit(string(item)) {
			continue
		} else {
			return false
		}
	}
	return true
}

func (args *Args)CheckType(key string,v interface{})bool{
	if !args.Check(){
		logrus.Println("GetValue false!")
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

func (args *Args)GetValue(key string)string{
	if !args.Check(){
		logrus.Println("GetValue false!")
		return ""
	}
	return args.CommandMapping[key]
}