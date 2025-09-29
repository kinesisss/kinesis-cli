package main
import (
	"github.com/kinesisss/kinesis-cli/internal/storageHandler"
	"fmt"
	"encoding/json"
	"time"
	"flag"
)

func AppendTaskToStorage(taskName string, creationTime time.Time, storagePath string) {
	StorageInstance, err := storageHandler.RetrieveTaskData(storagePath)
	if err != nil {
		fmt.Printf("--- we had an error retrieving the data stored in the file ---")
		fmt.Printf(err.Error())
	}
	dateKey := storageHandler.CreateDateKey(creationTime)
	userTask := storageHandler.UserTask{
		Id: len(StorageInstance.DailyStore[dateKey]),
		Name: taskName,
		State: "PENDING",
		CreatedAt: creationTime.String(),
	}
	StorageInstance.DailyStore[dateKey] = append(StorageInstance.DailyStore[dateKey], userTask)
	jbytes, jerr := json.Marshal(StorageInstance)
	if jerr != nil {
		fmt.Printf("... something went wrong when transforming your task \n")
	}
	storageHandler.WriteToStorageFile(storagePath, jbytes, 0666)	
}

func main(){
	task := flag.String("task", "<not-called>", "the task flag allows you to record a task")
	state := flag.Bool("state", false, "the state flag allows you to retrieve the current state of your tasks")
	done := flag.Int("done", -1, "the mod flag allows you to modify the state of a task")
	flag.Parse()
	storagePath := "./storage/tasks.json"
	if *task != "<not-called>"{
		timeNow := time.Now()
		canStore, _ := storageHandler.CheckStorage(storagePath)
		if !canStore{
			_, err := storageHandler.CreateStorage(storagePath)
			if err != nil {
				fmt.Printf("something happened while setting up your storage\n")
				fmt.Printf(err.Error())
			}
		}
 		AppendTaskToStorage(*task, timeNow, storagePath)			
	} 
	if *state{
		storageHandler.RenderTaskData(storagePath)
	}
    if *done != -1 {
		storageHandler.MarkStateAsDone(*done, storagePath)
	}
}