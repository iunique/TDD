package project2

import (
	"strings"
	"testing"
)

func TestArgs(test *testing.T){
	exp := []struct{
		flags string
		command string
		key byte
		real string
	}{
		{
			flags:"l:bool,d:int,f:string",
			command:"-l -d 8010 -f /usr/local",
			key:'1',
			real:"false",
		},
		{
			flags:"l:bool,d:int,f:string",
			command:"-l -d 8010 -f /usr/local",
			key:'d',
			real:"8080",
		},
		{
			flags:"l:bool,d:int,f:string",
			command:"-l -d 8010 -f /usr/local",
			key:'f',
			real:"/usr/local",
		},
	}
	for i,t :=range exp{
		args := new(Args)
		args.Init(t.flags,t.command)
		args.ParseFlags()
		args.ParseCommand()
		s:=args.GetValue(t.key)
		if(strings.Compare(s,t.real)!=0) {
			test.Errorf("第%d组出错，输入参数:flags:%v command:%v key:%v real:%s  得到结果:%s\n",i,t.flags,t.command,t.key,t.real,s)
		}
	}
}