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
		"26": lessons.ChannelsDirections,
		"27": lessons.Select,
		"28": lessons.Timeouts,
		"29": lessons.NonBlockingChannelOperations,
		"30": lessons.ClosingChannels,
		"31": lessons.RangeOverChannels,
		"32": lessons.Timers,
		"33": lessons.Tickers,
		"34": lessons.WorkerPools,
		"35": lessons.RateLimiting,
		"36": lessons.AtomicCounters,
		"37": lessons.Mutexes,
		"38": lessons.StatefulGoroutines,
		"39": lessons.Sorting,
		"40": lessons.SortingByFunctions,
		"41": lessons.Panic,
		"42": lessons.Defer,
		"43": lessons.CollectionFunctions,
		"44": lessons.StringFunctions,
		"45": lessons.StringFormatting,
		"46": lessons.RegularExpressions,
		"47": lessons.Json,
		"48": lessons.Time,
		"49": lessons.Epoch,
		"50": lessons.TimeFormattingParsing,
		"51": lessons.RandomNumbers,
		"52": lessons.NumberParsing,
		"53": lessons.URLParsing,
		"54": lessons.SHA1Hashes,
		"55": lessons.Base64Encoding,
		"56": lessons.ReadingFiles,
		"57": lessons.WritingFiles,
		"58": lessons.LineFilters,
		"59": lessons.CommandLineArguments,
		"60": lessons.CommandLineFlags,
		"61": lessons.EnvironmentVariables,
		"62": lessons.SpawningProcesses,
		"63": lessons.ExecProcesses,
		"64": lessons.Signals,
		"65": lessons.Exit,
	}

	//for k, v := range lessons {
	//	runLesson(k, v)
	//}

	k := "56"
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