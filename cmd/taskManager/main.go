package main
import (
	"github.com/kinesisss/kinesis-cli/internal/storageHandler"
	"fmt"
	"encoding/json"
	"time"
	"flag"
)

func CreateDateKey(targetTime time.Time) string{
	year, month, day := targetTime.Date()
	return fmt.Sprintf("%v-%v-%v", year, month, day)
}

func AppendTaskToStorage(taskName string, creationTime time.Time, storagePath string) {
	userTask := storageHandler.UserTask{
		Name: taskName,
		CreatedAt: creationTime.String(),
	}
	StorageInstance, err := storageHandler.RetrieveTaskData(storagePath)
	if err != nil {
		fmt.Printf("--- we had an error retrieving the data stored in the file ---")
		fmt.Printf(err.Error())
	}
	dateKey := CreateDateKey(creationTime)
	StorageInstance.DailyStore[dateKey] = append(StorageInstance.DailyStore[dateKey], userTask)
	jbytes, jerr := json.Marshal(StorageInstance)
	if jerr != nil {
		fmt.Printf("... something went wrong when transforming your task \n")
	}
	storageHandler.WriteToStorageFile(storagePath, jbytes, 0666)	
}

func main(){
	task := flag.String("task", "<not-called>", "the task flag allows you to record a flag")
	state := flag.Bool("state", false, "the state flag allows you to retrieve the current state of your tasks")
	flag.Parse()
	storagePath := "./storage/tasks.json"
	if *task != "<not-called>"{
		timeNow := time.Now()
		canStore, _ := storageHandler.CheckStorage(storagePath)
		if !canStore{
			fmt.Printf("... setting up your storage system\n")
			_, err := storageHandler.CreateStorage(storagePath)
			if err != nil {
				fmt.Printf("something happened while setting up your storage\n")
				fmt.Printf(err.Error())
			}
		}
 		AppendTaskToStorage(*task, timeNow, storagePath)			
	} 
	if *state{
		fmt.Printf("retrieving your tasks....\n")
		storageHandler.RetrieveTaskData(storagePath)
	}

}