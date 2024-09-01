package main

// This is a small go-based nice challenge where we need to mimic a single-threaded in-memory database.
// It should support a list of commands GET / SET / COUNT / DELETE in an interactive fashion.

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 04 November 2021

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// map to keep track of each unique value and its counters/total existence.
var valuesCounters map[int]int

// store all users data, this is the in-memory datastore.
var database map[string]int

// set of all valid commands with expected number of arguments.
var validCommands map[string]int

func init() {
	// initialize all maps.
	validCommands = make(map[string]int)
	database = make(map[string]int)
	valuesCounters = make(map[int]int)
}

// executeSETCommand adds a key/value pair.
func executeSETCommand(name string, value int) {
	oldValue := database[name]
	database[name] = value
	valuesCounters[value] = valuesCounters[value] + 1
	if _, exist := valuesCounters[oldValue]; exist {
		valuesCounters[oldValue] = valuesCounters[oldValue] - 1
	}
}

// executeGETCommand retreives value mapped to a name/key.
func executeGETCommand(name string) string {
	value, exist := database[name]
	if exist {
		// found name/key.
		return fmt.Sprintf("%d", value)
	}
	return "NULL"
}

// executeDELETECommand removes the pair from the data map.
func executeDELETECommand(name string) {
	value := database[name]
	delete(database, name)
	valuesCounters[value] = valuesCounters[value] - 1
}

// executeCOUNTCommand returns the number of keys having the specified value.
func executeCOUNTCommand(value int) int {
	count := 0
	for _, v := range database {
		if v == value {
			count += 1
		}
	}
	return count
}

func main() {

	var input string

	// fill the valid commands with the number of expected arguments included command itself.
	validCommands = map[string]int{"get": 2, "set": 3, "delete": 2, "count": 2}

	fmt.Println("\nWelcome on the interactive in-memory database [use CTRL-C to quit]")

	reader := bufio.NewScanner(os.Stdin)
	// loop to request a valid command.
	for {
		// prompt the user to enter the command.
		fmt.Printf("\n>> ")
		reader.Scan()
		// remove any space around and transform to uppercase.
		input = strings.ToLower(strings.TrimSpace(reader.Text()))

		if input == "" {
			// nothing entered so ignore.
			continue
		}
		// build a list of each word based on whitespace.
		args := strings.Fields(input)

		// check if command exist and retrieve its number of required args.
		argsNumber, valid := validCommands[args[0]]
		if !valid {
			// unknown command, so ignore and notify the user.
			fmt.Println("unknown command. valid commands are GET SET DELETE COUNT")
			continue
		}

		// user entered a valid command, so lets check the total number of arguments.
		if len(args) != argsNumber {
			// invalid syntax, so ignore and notify the user.
			fmt.Println("invalid syntax. find below the usage details.")
			continue
		}

		// process the command.
		switch args[0] {
		case "set":
			// convert the third option into integer (since value)
			// in case of failure, ignore the user request.
			if value, err := strconv.Atoi(args[2]); err == nil {
				executeSETCommand(args[1], value)
			} else {
				fmt.Println("invalid value. expects a number.")
			}

		case "get":
			fmt.Println(executeGETCommand(args[1]))
		case "delete":
			executeDELETECommand(args[1])
		case "count":
			// convert the third option from string to integer (since value)
			// in case of failure, ignore the user request.
			if value, err := strconv.Atoi(args[1]); err == nil {
				fmt.Println(executeCOUNTCommand(value))
			} else {
				fmt.Println("invalid value. expects a number.")
			}

		default:
			continue
		}
	}
}

/*
// Ouput:

~$ go run main.go

Welcome on the interactive in-memory database [use CTRL-C to quit]

>> set a 12

>> set b 12

>> get a
12

>> set c 12

>> get c
12

>> get d
NULL

>> count 12
3

>> delete a

>> get a
NULL

>> set b 13

>> get b
13

>>

*/
