package main

import (
	"fmt"
	"divo.net.pl/lessons"
)

type lesson_fn func()

func main() {
	runLesson("01", lessons.HelloWorld)
	runLesson("02", lessons.Values)
	runLesson("03", lessons.Variables)
	runLesson("04", lessons.Constants)
	runLesson("05", lessons.For)
	runLesson("06", lessons.IfElse)
	runLesson("07", lessons.Switch)
	runLesson("08", lessons.Arrays)
	runLesson("09", lessons.Slices)
	runLesson("10", lessons.Maps)
	runLesson("11", lessons.Range)
	runLesson("12", lessons.Functions)
	runLesson("13", lessons.MultipleReturnValues)
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