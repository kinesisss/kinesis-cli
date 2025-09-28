package storageHandler

import (
	"os"
	"fmt"
	"encoding/json"	
	"time"
)

type UserTask struct{
	Id int
	Name string
	State string
	CreatedAt string
}

func (u UserTask) RenderUserTask() string{
	return fmt.Sprintf("<%v> task created at <%v> \n", u.Name, u.CreatedAt)
}

func (u UserTask) ModifyState(state string) {
	u.State = state
}

type TaskStorage struct{
	DailyStore map[string][]UserTask
}

func CreateDateKey(targetTime time.Time) string{
	year, month, day := targetTime.Date()
	return fmt.Sprintf("%v-%v-%v", year, month, day)
}

func CheckStorage(dirName string) (bool, error){
	_, err := os.Stat(dirName)

	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateStorage(storageDir string) (os.File, error){
	rFilePtr, fileErr := os.Create(storageDir)
	if fileErr != nil{
		return *rFilePtr, fileErr
	}
	dummyTask := TaskStorage{
		DailyStore: map[string][]UserTask{},
	}
	jbytes, jerr := json.Marshal(dummyTask)
	if jerr != nil{
		fmt.Printf("something went wrong transforming your struct")
	}
	WriteToStorageFile(storageDir, jbytes, 0666)
	return *rFilePtr, nil
}

func WriteToStorageFile(storagePath string, data []byte, perm os.FileMode) {
	fmt.Printf("... persisting your task\n")
	werr := os.WriteFile(storagePath, data, perm)
	if werr != nil{
		fmt.Printf("something went wrong when persisting your file\n")
		fmt.Printf(werr.Error())
	}
	fmt.Printf("file has been persisted succesfully\n")
}

func RetrieveTaskData(storagePath string) (TaskStorage, error){
	data, err := os.ReadFile(storagePath)
	if err != nil{
		fmt.Printf("something went wrong when reading your file: \n")
		fmt.Printf(err.Error())
	}
	var StorageInstance TaskStorage
	jerr := json.Unmarshal(data, &StorageInstance)
	if jerr != nil{
		return StorageInstance, jerr
	}
	return StorageInstance, nil
}

func MarkStateAsDone(id int, storagePath string){
	TaskState, terr := RetrieveTaskData(storagePath)
	if terr != nil{
		fmt.Printf("We can't modify the state of your task")
		return
	}
	dateKey := CreateDateKey(time.Now())
	for i, uTask := range TaskState.DailyStore[dateKey]{
		if uTask.Id == id{
			TaskState.DailyStore[dateKey][i].State = "DONE"
			jbytes, _ := json.Marshal(TaskState)
			WriteToStorageFile(storagePath, jbytes, 0666)
		}
	}
	RenderTaskData(storagePath)
}

func RenderTaskData(storagePath string){
	TaskState, terr := RetrieveTaskData(storagePath)
	if terr != nil{
		fmt.Printf("We can't render your tasks")
		return
	}
	dateKey := CreateDateKey(time.Now())
	fmt.Printf("\nTask State [%v]:", dateKey)
	fmt.Printf("\n")
	for _, uTask := range TaskState.DailyStore[dateKey]{
		fmt.Printf("task: [%v] description: [%v] -- state: [%v]\n", uTask.Id, uTask.Name, uTask.State)
	}
	fmt.Printf("\n")
}