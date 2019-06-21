package service

import (
	"strconv"
	"strings"
	"testing"
)

func TestGame(test *testing.T){
	game := Init(3,5,"Fizz","Buzz")
	//str1:=game.str1
	str1,str2 :=game.str1,game.str2
	exp := []struct{
		real int
		result string
	}{
		{
			real:7,
			result:strconv.Itoa(7),
		},
		{
			real:12,
			result:str1,
		},
		{
			real:10,
			result:str2,
		},
		{
			real:15,
			result:str1+str2,
		},
		{
			real:9,
			result:str1,
		},
		{
			real:19,
			result:strconv.Itoa(19),
		},
		}
	for i,t :=range exp{
		var s string= game.Judge(t.real)
		if(strings.Compare(s,t.result)!=0) {
			test.Errorf("第%d组出错，输入参数:%v 得到结果:%v\n",i+1,t,s)
		}
	}
}