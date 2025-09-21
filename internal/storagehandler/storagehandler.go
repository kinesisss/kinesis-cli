package storagehandler

import (
	"os"
	"fmt"
)

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