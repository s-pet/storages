// nas
package storage

import (
	"errors"
	"log"
)

type nas struct {
	stgUuid string
	stgName string
	stgType string
	nasBase string
}

func (s nas) ExternalStorage_Init(attr StorageAttributes) error {
	e := validateName(attr["name"])
	if e != nil {
		log.Println(e.Error())
		return e
	}
	if attr["type"] != externalStorage_NAS {
		e = errors.New("Wrong StorageType " + attr["type"])
		log.Println(e.Error())
		return e
	}
	e = validateUuid(attr["uuid"])
	if e != nil {
		log.Println(e.Error())
		return e
	}
	s.stgUuid = attr["uuid"]
	s.stgName = attr["name"]
	s.stgType = attr["type"]
	s.nasBase = attr["base"]
	log.Println("Successfully created storage [" + s.stgName + ", " + s.stgType + ", " + s.stgUuid + ", " + s.nasBase + "]")
	return nil
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
