package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/thatisuday/commando"
)

func main() {

	// configure commando
	commando.
		SetExecutableName("todo").
		SetVersion("1.0.0").
		SetDescription("This tool lets you create a simple todo list.")

	// configure the root command
	commando.
		Register(nil).
		AddArgument("name", "name of the todo list", "todo").
		AddFlag("tasks,t", "comma separated list of tasks", commando.String, ""). // default ``
		AddFlag("location,l", "location to create file", commando.String, "./").  // default `./`
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `root` command...\n\n")
			name := args["name"].Value
			tasksStr, _ := flags["tasks"].GetString()
			location, _ := flags["location"].GetString()
			//split `tasks` into slice
			tasks := strings.Split(tasksStr, ",")
			//check if `location` ends in `/` if not append
			if !strings.HasSuffix(location, "/") {
				location = location + "/"
			}
			createTodo(name, tasks, location)

		})

	// configure info command
	commando.
		Register("info").
		SetShortDescription("This command lets you create a simple todo list").
		AddArgument("name", "name of the todo list", "todo.txt").
		AddFlag("tasks,t", "comma separated list of tasks", commando.String, nil).
		AddFlag("location,l", "location to create file", commando.String, "./").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `info` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// parse command-line arguments
	commando.Parse(nil)

}

func createTodo(name string, tasks []string, location string) {
	if !strings.Contains(name, ".txt") {
		f, err := os.Create(location + name + ".txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		f.WriteString(name + "\n\n")
		for _, task := range tasks {
			f.WriteString("[ ] " + task + "\n")
		}
	} else {
		f, err := os.Create(location + name)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.WriteString(name + "\n\n")
		for _, task := range tasks {
			f.WriteString("[ ] " + task + "\n")
		}
	}

}
