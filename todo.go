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
		AddFlag("tasks,t", "comma separated list of tasks", commando.String, nil). // default ``
		AddFlag("sub-tasks,s", "a comma separated list of sub-tasks for a task, can only be used on a single task", commando.String, "nil").
		AddFlag("location,l", "location to create file", commando.String, "./"). // default `./`
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			name := args["name"].Value
			tasksStr, _ := flags["tasks"].GetString()
			subTaskStr, _ := flags["sub-tasks"].GetString()
			location, _ := flags["location"].GetString()
			//split `tasks` into slice
			tasks := strings.Split(tasksStr, ",")
			subTasks := strings.Split(subTaskStr, ",")
			//check if `location` ends in `/` if not append
			if !strings.HasSuffix(location, "/") {
				location = location + "/"
			}
			createTodo(&name, &tasks, &subTasks, &location)
			// openTodoInVSCode(name)
		})

	// configure info command
	commando.
		Register("info").
		SetShortDescription("This command lets you create a simple todo list").
		AddArgument("name", "name of the todo list", "todo.md").
		AddFlag("tasks,t", "comma separated list of tasks", commando.String, nil).
		AddFlag("sub-tasks,s", "a comma separated list of sub-tasks for a task, can only be used on a single task", commando.String, "nil").
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

func createTodo(name *string, tasks *[]string, subTasks *[]string, location *string) {
	removeEmpty(tasks)
	removeEmpty(subTasks)

	n := *name
	t := *tasks
	s := *subTasks
	l := *location

	if len(t) > 1 && len(s) > 0 {
		fmt.Println("Cant have multiple tasks and subtasks")
	} else if !strings.Contains(n, ".md") {
		f, err := os.Create(l + n + ".md")
		if err != nil {
			fmt.Println(err)
			return
		}
		f.WriteString("# ✅✅ " + n + " ✅✅\n\n")
		for _, task := range t {
			if len(s) == 0 {
				//no sub tasks so make the tasks the check list
				f.WriteString("- [ ] " + strings.TrimSpace(task) + "\n")
			} else {
				f.WriteString("### " + strings.TrimSpace(task) + "\n")

				for _, subTask := range s {
					f.WriteString("- [ ] " + strings.TrimSpace(subTask) + "\n")
				}
			}
		}
	}

	// else {
	// 	f, err := os.Create(l + n)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	f.WriteString(n + "\n\n")
	// 	for _, task := range t {
	// 		f.WriteString("[ ] " + strings.TrimSpace(task) + "\n")
	// 	}
	// }

}

func openTodoInVSCode(name string) {
	if !strings.Contains(name, ".md") {
		exec.Command("code", name+".md").Output()
	} else {
		exec.Command("code", name).Output()
	}
}

func removeEmpty(slice *[]string) {
	i := 0
	p := *slice
	for _, entry := range p {
		if strings.Trim(entry, " ") != "" {
			p[i] = entry
			i++
		}
	}
	*slice = p[0:i]
}
