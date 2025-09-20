package main
import (
	"github.com/kinesisss/kinesis-cli/internal/cliparser"
	"github.com/kinesisss/kinesis-cli/internal/storagehandler"
	"fmt"
	"time"
)

type UserTask struct{
	name string
	createdAt string
}

func createTask(taskName string, creationTime time.Time) UserTask{
	userTask := UserTask{
		name: taskName,
		createdAt: creationTime.String(),
	}
	return userTask
}

func (u UserTask) renderUserTask() string{
	return fmt.Sprintf("<%v> task created at <%v> \n", u.name, u.createdAt)
}

func main(){
	taskName := cliparser.TaskFlagParser()
	timeNow := time.Now()
	currentTask := createTask(taskName, timeNow)
	fmt.Printf(currentTask.renderUserTask())
	canStore, _ := storagehandler.CheckStorage(".", "storage")
	if canStore{
		fmt.Printf("we can definetely store your data\n")
	}

}