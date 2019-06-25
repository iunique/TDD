package project2

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

func (args *Args)IntParse(key string,v string)bool{
	defer func() {
		recover()
	}()
	_,err:=strconv.Atoi(v)
	if(err!=nil){
		args.CommandMapping[key]=args.DefaultMapping[key]//设置默认值
		logrus.Println("IntParse Fail!")
	}else {
		args.CommandMapping[key]=v
		return true
	}
	return false
}
