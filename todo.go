package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Command struct {
	name        string
	description string
	call        func([]string)
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
	todos := []ToDo{}
	commands = []Command{
		{
			name:        "help",
			description: "Provides detail on a specific command",
			call: func(params []string) {
				fmt.Printf("this is a help function for %v", params[0])
			},
		},
		{
			name:        "add",
			description: "Adds a new items to the list",
			call: func(params []string) {
				fmt.Printf("adding %v to todo list", params[1])

				if len(params) < 3 {
					log.Fatalln("Not enough parameters provided (2)")
					return
				}
				if len(params) > 5 {
					log.Fatalln("Too many parameters provided (2)")
					return
				}

				title := params[1]
				desc := params[2]
				dueDateTime, _ := time.Parse("yyyy/MM/dd", "0000/00/00")
				comDateTime, _ := time.Parse("yyyy/MM/dd", "0000/00/00")

				if len(params) > 3 {
					dueDate, dDErr := time.Parse("yyyy/MM/dd", params[3])
					dueTime, dTErr := time.Parse("HH:mm:ss", params[3])

					if dDErr != nil && dTErr != nil {
						log.Fatalln(dDErr)
						log.Fatalln(dTErr)
						return
					}
					dueDateTime = time.Date(dueDate.Year(), dueDate.Month(), dueDate.Day(), dueTime.Hour(), dueTime.Minute(), dueTime.Second(), dueTime.Nanosecond(), dueTime.Location())
				}

				if len(params) > 4 {
					comDate, dDErr := time.Parse("yyyy/MM/dd", params[3])
					comTime, dTErr := time.Parse("HH:mm:ss", params[3])

					if dDErr != nil && dTErr != nil {
						log.Fatalln(dDErr)
						log.Fatalln(dTErr)
						return
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
				fmt.Printf("Added item to todo list (%v)", len(todos))
			},
		},
	}

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	split := strings.Split(text, " ")
	fmt.Printf("SPLIT: %v\n", split)
	for k, v := range commands {
		fmt.Printf("%v - input: %v; cmd.Name: %v\n", k, split[0], v.name)
		if v.name == split[0] {
			v.call(split)
		}
	}

	fmt.Println(text)
	main()
}
