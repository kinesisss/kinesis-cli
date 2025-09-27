package main
import (
	"github.com/kinesisss/kinesis-cli/internal/storageHandler"
	"fmt"
	"encoding/json"
	"time"
	"flag"
)

func createTask(taskName string, creationTime time.Time) storageHandler.UserTask{
	userTask := storageHandler.UserTask{
		Name: taskName,
		CreatedAt: creationTime.String(),
	}
	return userTask
}

func main(){
	task := flag.String("task", "<not-called>", "the task flag allows you to record a flag")
	state := flag.Bool("state", false, "the state flag allows you to retrieve the current state of your tasks")
	flag.Parse()
	storagePath := "./storage/tasks.json"
	if *task != "<not-called>"{
		timeNow := time.Now()
		currentTask := createTask(*task, timeNow)		
		fmt.Printf(currentTask.RenderUserTask())
		canStore, _ := storageHandler.CheckStorage(storagePath)
		jbytes, jerr := json.Marshal(currentTask)
		if jerr != nil {
			fmt.Printf("... something went transforming your task\n")
		}
		if !canStore{
			fmt.Printf("... setting up your storage system\n")
			_, err := storageHandler.CreateStorage(storagePath)
			if err != nil {
				fmt.Printf("something happened while setting up your storage\n")
				fmt.Printf(err.Error())
			}
		} 
		storageHandler.WriteToStorageFile(storagePath, jbytes, 0666)	
	} 
	if *state{
		fmt.Printf("retrieving your tasks....\n")
		storageHandler.RetrieveTaskData(storagePath)
	}

}