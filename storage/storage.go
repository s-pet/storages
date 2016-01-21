// storage
package storage

import (
	"errors"
	"log"
)

const (
	externalStorage_NAS                = "NAS"
	externalStorage_GoogleGloudStorage = "Google Cloud Storage"
	externalStorage_AmazonS3           = "Amazon S3"
)

type ExternalStorage interface {
	ExternalStorage_GetName() string
	ExternalStorage_GetType() string
	ExternalStorage_GetUuid() string
	ExternalStorage_Store(doc string) error
}

type StorageAttributes map[string]string

func CreateStorage(attr StorageAttributes) (ExternalStorage, error) {
	var (
		s ExternalStorage
		e error
	)
	switch attr["type"] {
	case externalStorage_NAS:
		s, e = nas{}.construct(attr)
	default:
		e = errors.New("StorageType " + attr["type"] + " is not supported")
		log.Println(e.Error())
	}
	return s, e
}

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
