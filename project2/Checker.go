package project2

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

func (args *Args)CheckCommand(s string)bool{//分析字符串是否为Command以及Command是否合法
	if(len(s)==2&&s[0]=='-'&&s[1]>='a'&&s[1]<='z'){
		_,isExist:=args.FlagsMapping[s[1:]]
		if(!isExist){
			logrus.Println("不存在命令")
		}
		return isExist
	}
	logrus.Println("该字串不为命令")
	return false
}


//判断初始化
func (args *Args)CheckInit()bool{
	if args.Flags==""||args.Command==""||args.FlagsMapping==nil||args.CommandMapping==nil{
		logrus.Errorf("the args is not ready")
		return false
	}
	return true
}

//判断类型
func (args *Args)CheckType(key string,v interface{})bool{
	if !args.CheckInit(){
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


