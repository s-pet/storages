// storage-attributes
package storage

import (
	"errors"
)


type StorageAttributes map[string]string


func validateName(stgName string) error {
	if len(stgName) == 0 {
		return errors.New("Storage name cannot be empty")
	}
	return nil
}

func validateUuid(stgUuid string) error {
	if len(stgUuid) == 0 {
		return errors.New("Storage uuid cannot be empty")
	}
	return nil
}



