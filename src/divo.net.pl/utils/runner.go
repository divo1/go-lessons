package utils

import (
	"divo.net.pl/lessons"
	"fmt"
	"strings"
)

func Runner(moduleToRun string, whichRun []string) {
	modules := make(map[string]func())
	module := strings.ToLower(moduleToRun)
	switch {
	case module == "lessons":
		modules = lessons.Lessons()
	}

	if len(modules) == 0 {
		panic("Bad module name")
	}

	for _, v := range whichRun {
		runLesson(v, modules[v])
	}
}

func runLesson(number string, call func()) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("~  Lesson number %-12s", number)
	fmt.Printf("~\n")
	fmt.Println("~  ------------------------  ~")
	call()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("")
}
