package main
import (
	"github.com/kinesisss/kinesis-cli/internal/cliparser"
	"github.com/kinesisss/kinesis-cli/internal/storagehandler"
	"fmt"
	"encoding/json"
	"time"
)

type UserTask struct{
	Name string
	CreatedAt string
}

func createTask(taskName string, creationTime time.Time) UserTask{
	userTask := UserTask{
		Name: taskName,
		CreatedAt: creationTime.String(),
	}
	return userTask
}

func (u UserTask) renderUserTask() string{
	return fmt.Sprintf("<%v> task created at <%v> \n", u.Name, u.CreatedAt)
}

func main(){
	taskName := cliparser.TaskFlagParser()
	timeNow := time.Now()
	currentTask := createTask(taskName, timeNow)
	storagePath := "./storage/tasks.json" 
	fmt.Printf(currentTask.renderUserTask())
	canStore, _ := storagehandler.CheckStorage(storagePath)
	jbytes, jerr := json.Marshal(currentTask)
	if jerr != nil {
		fmt.Printf("... something went transforming your task")
	}
	if !canStore{
		fmt.Printf("... setting up your storage system")
		_, err := storagehandler.CreateStorage(storagePath)
		if err != nil {
			fmt.Printf("something happened while setting up your storage")
			fmt.Printf(err.Error())
		}
	} 
	storagehandler.WriteToStorageFile(storagePath, jbytes, 0666)

}