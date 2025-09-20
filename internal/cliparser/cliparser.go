package cliparser
import "flag"

func TaskFlagParser() string{
	task_name := flag.String("task", "<name-of-your-task>", "the task flag allows you to record a flag")
	flag.Parse()
	return *task_name
}