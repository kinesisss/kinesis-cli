package main
import (
	"github.com/kinesisss/kinesis-cli/internal/storagehandler"
	"fmt"
	"encoding/json"
	"time"
	"flag"
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
	task := flag.String("task", "<not-called>", "the task flag allows you to record a flag")
	state := flag.Bool("state", false, "the state flag allows you to retrieve the current state of your tasks")
	flag.Parse()
	if *task != "<not-called>"{
		timeNow := time.Now()
		currentTask := createTask(*task, timeNow)
		storagePath := "./storage/tasks.json" 
		fmt.Printf(currentTask.renderUserTask())
		canStore, _ := storagehandler.CheckStorage(storagePath)
		jbytes, jerr := json.Marshal(currentTask)
		if jerr != nil {
			fmt.Printf("... something went transforming your task\n")
		}
		if !canStore{
			fmt.Printf("... setting up your storage system\n")
			_, err := storagehandler.CreateStorage(storagePath)
			if err != nil {
				fmt.Printf("something happened while setting up your storage\n")
				fmt.Printf(err.Error())
			}
		} 
		storagehandler.WriteToStorageFile(storagePath, jbytes, 0666)	
	} 
	if *state{
		fmt.Printf("retrieving your tasks....\n")
	}

}