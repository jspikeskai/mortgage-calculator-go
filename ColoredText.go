package main

const RESET = "\u001B[0m"
const RED = "\u001B[31m"
const GREEN = "\u001B[32m"
const YELLOW = "\u001B[33m"
const BLUE = "\u001B[34m"
const CYAN = "\u001B[36m"

func Color(colorCode string, text string) string {
	return colorCode + text + RESET
}
