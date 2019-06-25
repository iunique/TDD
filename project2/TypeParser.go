package project2

import (
	"github.com/sirupsen/logrus"
)

func (args *Args)TypeParse(tmpType string,key string,v string)bool{
	switch tmpType {
	case "string":
		if (args.StringParse(key,v)) {
			logrus.Println("string分解成功")
			return true
		}
	case "int":
		if (args.IntParse(key,v)) {
			logrus.Println("int分解成功")
			return true
		}
	case "bool":
		if (args.BoolParse(key,v)) {
			logrus.Println("nool分解成功")
			return true
		}
	}
	return false
}
