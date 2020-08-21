package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/thatisuday/commando"
)

func main() {

	// configure commando
	commando.
		SetExecutableName("todo").
		SetVersion("1.0.0").
		SetDescription("This tool lets you create a simple todo list. Opens using `code` so make sure you have that set up 😉")

	// configure the root command
	commando.
		Register(nil).
		AddArgument("name", "name of the todo list", "todo").
		AddFlag("tasks,t", "comma separated list of tasks", commando.String, ""). // default ``
		AddFlag("location,l", "location to create file", commando.String, "./").  // default `./`
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
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
			openTodoInVSCode(name)
		})

	// configure info command
	commando.
		Register("info").
		SetShortDescription("This command lets you create a simple todo list").
		AddArgument("name", "name of the todo list", "todo.txt").
		AddFlag("tasks,t", "comma separated list of tasks", commando.String, nil).
		AddFlag("location,l", "location to create file", commando.String, "./").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
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
		f.WriteString("✅✅ " + name + " ✅✅\n\n")
		for _, task := range tasks {
			f.WriteString("[ ] " + strings.TrimSpace(task) + "\n")
		}
	} else {
		f, err := os.Create(location + name)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.WriteString(name + "\n\n")
		for _, task := range tasks {
			f.WriteString("[ ] " + strings.TrimSpace(task) + "\n")
		}
	}

}

func openTodoInVSCode(name string) {
	if !strings.Contains(name, ".txt") {
		exec.Command("code", name+".txt").Output()
	} else {
		exec.Command("code", name).Output()
	}
}
