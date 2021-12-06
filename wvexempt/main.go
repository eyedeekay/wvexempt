package main

import (
	"github.com/wvexempt/wvexempt"

func main(){
	err := wvexempt.Exempt()
	if err != nil{
		panic(err)
	}

}