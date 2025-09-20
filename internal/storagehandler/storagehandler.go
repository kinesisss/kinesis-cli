package storagehandler

import (
	"os"
)

func CheckStorage(dirName string, nameToCheck string) (bool, error){
	files, err := os.ReadDir(dirName)

	if err != nil {
		return false, err
	}

	for _, file := range files{
		if file.Name() == nameToCheck{
			return true, nil
		}
	}
	return false, nil
}