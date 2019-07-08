package project2

import (
	"strconv"
)

func (args *Args) BoolParse(key string, v string) bool {
	defer func() {
		recover()
	}()
	_, err := strconv.ParseBool(v)
	if err != nil {
		args.CommandMapping[key] = args.DefaultMapping[key] //设置默认值
		//logrus.Println("BoolParse Fail!")
	} else {
		args.CommandMapping[key] = v
		return true
	}
	return false
}
