package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Command struct {
	name        string
	help        string
	description string
	call        func([]string) int
}

type ToDo struct {
	title          string
	description    string
	dueDate        time.Time
	completionDate time.Time
	creationDate   time.Time
}

var commands []Command
var todos []ToDo

func main() {
	setup()
	loop()
}

func loop() {
	fmt.Println("loop start")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	split := strings.Split(text, " ")
	errorCode := 0

	cmd := getCmd(split[0])
	if cmd != nil {
		errorCode = cmd(split)
	}

	if errorCode > 0 {
		fmt.Printf("Error code: %v", errorCode)
		return
	}

	reader.Reset(os.Stdin)
	reader = nil
	loop()
}

func getCmd(input string) func([]string) int {
	for _, cmd := range commands {
		if strings.ToLower(cmd.name) == strings.ToLower(input) {
			return cmd.call
		}
	}
	for _, cmd := range commands {
		fmt.Printf("\t%v - %v\n", cmd.name, cmd.description)
	}
	return nil
}

func setup() {
	todos := []ToDo{}
	commands = []Command{
		{
			name:        "help",
			help:        "Please provide a paramater to search the help documentation\nFor example: `help commands`",
			description: "Provides detail on a specific command",
			call: func(params []string) int {
				fmt.Printf("Params Provided: %v\n", params)
				fn := ""
				if len(params) > 1 {
					fn = params[1]
				}
				if fn == "" {
					fmt.Println(commands[0].help)
				} else {
					for _, v := range commands {
						if strings.ToLower(v.name) == strings.ToLower(fn) {
							fmt.Println(v.help)
							return 0
						}
					}
				}
				fmt.Printf("%v\n", params[0])
				return 1
			},
		},
		{
			name:        "add",
			help:        "Takes three paramters:\n title: string - The name of your to list item\n desc: A description for the item\n dueDateTime?: Sets the due date (optional)",
			description: "Adds a new items to the list",
			call: func(params []string) int {

				if len(params) < 3 {
					fmt.Println("Not enough parameters provided (a, b)")
					return 1
				}
				if len(params) > 5 {
					fmt.Println("Too many parameters provided (a, b)")
					return 2
				}

				fmt.Printf("adding %v to todo list\n", params[1])

				title := params[1]
				desc := params[2]
				dueDateTime, _ := time.Parse("yyyy/MM/dd", "0000/00/00")
				comDateTime, _ := time.Parse("yyyy/MM/dd", "0000/00/00")

				if len(params) > 3 {
					dueDate, dDErr := time.Parse("yyyy/MM/dd", params[3])
					dueTime, dTErr := time.Parse("HH:mm:ss", params[3])

					if dDErr != nil && dTErr != nil {
						fmt.Println(dDErr)
						fmt.Println(dTErr)
						return 3
					}
					dueDateTime = time.Date(dueDate.Year(), dueDate.Month(), dueDate.Day(), dueTime.Hour(), dueTime.Minute(), dueTime.Second(), dueTime.Nanosecond(), dueTime.Location())
				}

				if len(params) > 4 {
					comDate, dDErr := time.Parse("yyyy/MM/dd", params[3])
					comTime, dTErr := time.Parse("HH:mm:ss", params[3])

					if dDErr != nil && dTErr != nil {
						fmt.Println(dDErr)
						fmt.Println(dTErr)
						return 4
					}
					comDateTime = time.Date(comDate.Year(), comDate.Month(), comDate.Day(), comTime.Hour(), comTime.Minute(), comTime.Second(), comTime.Nanosecond(), comTime.Location())
				}

				todo := ToDo{
					title:          title,
					description:    desc,
					dueDate:        dueDateTime,
					completionDate: comDateTime,
					creationDate:   time.Now(),
				}

				todos = append(todos, todo)
				fmt.Printf("Added item to todo list (%v)\n", len(todos))
				return 0
			},
		},
	}
}
