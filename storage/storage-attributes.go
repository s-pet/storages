// storage-attributes
package storage

import (
	"errors"
)


type StorageAttributes map[string]string

type stgAttrType struct {
	stgType, stgUuid, stgName string
}

func constructStgAttr(sType, sUuid, sName string) (stgAttrType, error) {
	var (
		stgAttr stgAttrType
		err error
	)
	if len(sType) == 0 {
		err = errors.New("Storage Type must not be empty: " + sType + ", " + sUuid + "," + sName)
	}
	if len(sUuid) == 0 {
		err = errors.New("Storage Uuid must not be empty: " + sType + ", " + sUuid + "," + sName)
	}
	if len(sName) == 0 {
		err = errors.New("Storage Name must not be empty: " + sType + ", " + sUuid + "," + sName)
	}
	if err == nil {
		stgAttr = stgAttrType{stgType: sType, stgUuid: sUuid, stgName: sName}
	}
	return stgAttr, err
}



