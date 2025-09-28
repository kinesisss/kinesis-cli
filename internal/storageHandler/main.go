package storageHandler

import (
	"os"
	"fmt"
	"encoding/json"	
)

type UserTask struct{
	Name string
	CreatedAt string
}

type TaskStorage struct{
	DailyStore map[string][]UserTask
}

func (u UserTask) RenderUserTask() string{
	return fmt.Sprintf("<%v> task created at <%v> \n", u.Name, u.CreatedAt)
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