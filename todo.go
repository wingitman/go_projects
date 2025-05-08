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
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	split := strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\n", "|"), " ", "|"), "|")
	errorCode := 0

	for i, s := range split {
		fmt.Printf("%v:%v\n",i,s)
		if i == 0 {
			continue
		}
		for _, v := range commands {
			fmt.Printf("%v, %v\n", s, v.name)
			if strings.ToLower(v.name) == strings.ToLower(s) {
				fmt.Printf("%v - invoking %v\n", i, v.name)
				errorCode = v.call(split)
			}
		}
	}

	if errorCode > 0 {
		fmt.Println("Invalid command, please choose one of the following commands:")
		for _, v := range commands {
			fmt.Printf("%v - %v\n", v.name, v.description)
		}
		fmt.Println()
	}
	loop()
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












