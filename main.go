package main

import (
	questions "challenges/questions/leetcode"
)

// designpatterns "challenges/design-patterns"

func main() {
	// ALGORITHMS
	// algorithms.RunBubbleSort()
	// algorithms.BubbleSort()
	// numbers := []int{212, 18, 110, 454, 34, 20, 110}
	// target := 18
	// search.BinarySearch(numbers, target)
	// // search.RunPreOrderBinarySearch()
	// datastructure.RunQueue()

	// --- Practice Questions ---

	// -- ADVENT OF CODE --
	/*
		2021
	*/
	// adventofcode2021.Submarine() day one
	// adventofcode2021.Dive() // day two
	// adventofcode2021.DivePartTwo() // day two
	/*
		2023
	*/
	// adventofcode2023.Trebuchet() // day one
	/*
		2024
	*/
	// adventofcode2024.DayOneHistorianHysteria()
	// adventofcode2024.DayTwo()

	// -- Leet Code --
	leetCode := questions.NewLeetCode()
	leetCode.Run()

	// CONCURRENCY
	// concurrency.TreasureHunt()
	// concurrency.SimpleExample()
	// concurrency.RunCrawler()
	// concurrency.FileTextSearch()
	// concurrency.AdvancedConcurrency()
	// concurrency.RunRequestThrottle()

	// TOOLS
	// filestream.RunFileStream()

	// EXPERIMENTS
	// experiments.RunRobotBuilder()

	// DESIGN PATTERNS
	// designpatterns.SingleResponsibilityPrinciple()
	// designpatterns.DependencyInversionPrinciple()
	// designpatterns.TestFactoryGenerator()
	// designpatterns.TestStrategyPattern()

	// PRACTICE
	// practice.Run()
}
