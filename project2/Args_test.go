package project2

import (
	"strings"
	"testing"
)

func TestArgs(test *testing.T) {
	exp := []struct {
		flags   string
		command string
		key     string
		real    interface{}
	}{
		{
			flags:   "l:bool,d:int,f:string",
			command: "-l -d 8010 -f /usr/local",
			key:     "l",
			real:    false,
		},
		{
			flags:   "l:bool,d:int,f:string",
			command: "-l -d 8010 -f /usr/local",
			key:     "d",
			real:    8080,
		},
		{
			flags:   "l:bool,d:int,f:string",
			command: "-l -d 8010 -f /usr/local",
			key:     "f",
			real:    "/usr/local",
		},
		{
			flags:   "l:bool,d:int,f:string,s:string",
			command: "-d -9 -f /usr/local -s",
			key:     "l",
			real:    false,
		},
		{
			flags:   "l:bool,d:int,f:string,s:string",
			command: "-d -9 -f /usr/local -s",
			key:     "f",
			real:    "/usr/local",
		},
		{
			flags:   "l:bool,d:int,f:string,s:string",
			command: "-d -9 -f /usr/local -s",
			key:     "s",
			real:    "",
		},
		{
			flags:   "l:bool,d:int,f:string",
			command: "-d 8080",
			key:     "l",
			real:    false,
		},
		{
			flags:   "l:bool,d:int,f:string",
			command: "-d 8080",
			key:     "d",
			real:    8080,
		},
		{
			flags:   "l:bool,d:int,f:string",
			command: "-d 8080",
			key:     "f",
			real:    ".",
		},
		{
			flags:   "l:bool,g:int,s:string",
			command: "-l -s test -g 80",
			key:     "l",
			real:    false,
		},
		{
			flags:   "l:bool,g:int,s:string",
			command: "-l -s test -g 80",
			key:     "s",
			real:    "test",
		},
		{
			flags:   "l:bool,g:int,s:string",
			command: "-l -s test -g 80",
			key:     "g",
			real:    80,
		},
	}
	for i, t := range exp {
		args := new(Args)
		args.Init(t.flags, t.command)
		args.DefaultMapping["bool"] = "false"
		args.DefaultMapping["int"] = "0"
		args.DefaultMapping["string"] = "default"
		args.ParseFlags()
		args.ParseCommand()
		s := args.GetValue(t.key)
		if !args.CheckType(t.key, t.real) {
			test.Errorf("第%d组类型错误，输入参数:flags:%v command:%v key:%v real:%v  得到结果:%s\n", i, t.flags, t.command, t.key, t.real, s)
		}
		if strings.Compare(args.GetValue(t.key), s) != 0 {
			test.Errorf("第%d组结果错误，输入参数:flags:%v command:%v key:%v real:%v  得到结果:%s\n", i, t.flags, t.command, t.key, t.real, s)
		}
	}
}
