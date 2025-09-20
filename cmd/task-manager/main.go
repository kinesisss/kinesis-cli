package main
import (
	"fmt"
	"github.com/kinesisss/kinesis-cli/internal/cliparser"
)

func main(){
	task_name := cliparser.TaskFlagParser()
	fmt.Printf("you have added <%s> to your tasks\n", string(task_name))
}