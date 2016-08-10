package main

import (
	"fmt"
	"divo.net.pl/lessons"
)

type lesson_fn func()

func main() {
	lessons := map[string]lesson_fn {
		"01": lessons.HelloWorld,
		"02": lessons.Values,
		"03": lessons.Variables,
		"04": lessons.Constants,
		"05": lessons.For,
		"06": lessons.IfElse,
		"07": lessons.Switch,
		"08": lessons.Arrays,
		"09": lessons.Slices,
		"10": lessons.Maps,
		"11": lessons.Range,
		"12": lessons.Functions,
		"13": lessons.MultipleReturnValues,
		"14": lessons.VariadicFunctions,
		"15": lessons.Closures,
		"16": lessons.Recursion,
		"17": lessons.Pointers,
		"18": lessons.Structs,
		"19": lessons.Methods,
		"20": lessons.Interfaces,
		"21": lessons.Errors,
		"22": lessons.Goroutines,
		"23": lessons.Channels,
		"24": lessons.ChannelsBuffering,
		"25": lessons.ChannelsSynchronization,
	}
	
	//for k, v := range lessons {
	//	runLesson(k, v)
	//}
	
	k := "25"
	runLesson(k, lessons[k])
}

func runLesson(number string, call lesson_fn) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("~  Lesson number %-12s", number)
	fmt.Printf("~\n")
	fmt.Println("~  ------------------------  ~")
	call()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("")
}