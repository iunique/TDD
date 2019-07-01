package project3

import (
	"fmt"
	"testing"
)

func TestArgs(test *testing.T){
	exp := []struct{
		flags string
		command string
		key string
		real interface{}
	}{
		{
			flags:"l:bool,g:int,s:string",
			command: "-l -s test -g 80",
			key:"l",
			real:false,
		},
		{
			flags:"l:bool,g:int,s:string",
			command: "-l -s test -g 80",
			key:"s",
			real:"test",
		},
		/*
		{
			flags:"l:bool,g:int,s:string",
			command: "-l -s test -g 80",
			key:"g",
			real:80,
		},*/
	}
	for _,t := range exp{
		args := new(Args)
		args.Init(t.flags, t.command )
		args.ParseCommand()
		result := args.GetValue(t.key)
		fmt.Println(result)
		if !AssertEqual(t.real,result) {
			test.Error("Error")
		}
	}
}

func AssertEqual(f1 interface{},f2 interface{})bool{
	return f1 == f2
}
