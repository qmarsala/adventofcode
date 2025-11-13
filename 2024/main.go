package main

import (
	"2024/day01"
	"2024/day02"
	"2024/day03"
	"2024/day04"
	"2024/day05"
	"2024/day06"
	"2024/day07"
	"fmt"
	"time"
)

type Day struct {
	RunDay func(string)
	Name   string
}

func main() {
	fmt.Print("Merry Christmas!\n\n")
	days := []Day{
		{RunDay: day01.RunDay, Name: "Day 01"},
		{RunDay: day02.RunDay, Name: "Day 02"},
		{RunDay: day03.RunDay, Name: "Day 03"},
		{RunDay: day04.RunDay, Name: "Day 04"},
		{RunDay: day05.RunDay, Name: "Day 05"},
		{RunDay: day06.RunDay, Name: "Day 06"},
		{RunDay: day07.RunDay, Name: "Day 07"},
	}
	profile("All Days", func() {
		for _, d := range days {
			profile(d.Name, func() {
				d.RunDay(d.Name)
			})
		}
	})
}

func profile(name string, op func()) {
	start := time.Now()
	op()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("%v completed in %v\n---\n", name, elapsed)
}
