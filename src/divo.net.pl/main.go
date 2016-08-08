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