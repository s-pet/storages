// interface for external storages
package storage

import (
	"errors"
)

const (
	externalStorage_NAS = "NAS" //Network Attached Storage; e.g. a local directory
	externalStorage_GCS = "GCS" //Google Cloud Storage (FOR FUTURE USE)
	externalStorage_AS3 = "AS3" //Amazon S3 (FOR FUTURE USE)
)

type ExternalStorage interface {
	ExternalStorage_Init(attr StorageAttributes) error
	ExternalStorage_GetName() string
	ExternalStorage_GetType() string
	ExternalStorage_GetUuid() string
	ExternalStorage_Store(doc string) error
}


func InitExternalStorage(attr StorageAttributes) (ExternalStorage, error) {
	var (
		stg     ExternalStorage
		err     error
	)
	
	if err != nil {
		switch attr["type"] {
			case externalStorage_NAS:
	     		stg = nas{}
			default:
				err = errors.New("StorageType \"" + attr["type"] + "\" is not supported")
		}
	}
	
	if err == nil {
		stg.ExternalStorage_Init(attr)
	}
	return stg, err
}
