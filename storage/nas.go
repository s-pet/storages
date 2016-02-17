// nas
package storage

import (
	"log"
)

type nas struct {
	stgUuid string
	stgName string
	stgType string
	nasBase string
}

func (s nas) ExternalStorage_Init(attr StorageAttributes) (ExternalStorage, error) {
	s.stgUuid = attr["uuid"]
	s.stgName = attr["name"]
	s.stgType = attr["type"]
	s.nasBase = attr["base"]
	log.Println("Successfully initialized nas storage [" + s.stgName + ", " + s.stgType + ", " + s.stgUuid + ", " + s.nasBase + "]")
	return s, nil
}

func (s nas) ExternalStorage_GetName() string {
	return s.stgName
}

func (s nas) ExternalStorage_GetType() string {
	return s.stgType
}

func (s nas) ExternalStorage_GetUuid() string {
	return s.stgUuid
}

func (s nas) ExternalStorage_Store(doc string) error {
	log.Println("Sucessfully stored " + doc)
	return nil
}
