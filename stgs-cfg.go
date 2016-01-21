// storageConfig
package main

import (
	"encoding/xml"
	"github.com/s-pet/storages/storage"
	"io/ioutil"
	"log"
	"os"
)

type TAGstorages struct {
	XMLName     xml.Name     `xml:"storages"`
	StorageList []TAGstorage `xml:"storage"`
}

type TAGstorage struct {
	XMLName       xml.Name       `xml:"storage"`
	AttributeList []TAGattribute `xml:"attribute"`
}

type TAGattribute struct {
	XMLName xml.Name `xml:"attribute"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:"value,attr"`
}

func readStorageConfig(xmlFilename string) ([]storage.StorageAttributes, error) {
	stgAttrList := make([]storage.StorageAttributes, 0)

	xmlFile, err := os.Open(xmlFilename)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer xmlFile.Close()

	b, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var tagStorages TAGstorages
	xml.Unmarshal(b, &tagStorages)

	for _, stg := range tagStorages.StorageList {
		stgAttr := make(storage.StorageAttributes)
		for _, attr := range stg.AttributeList {
			stgAttr[attr.Key] = attr.Value
		}
		stgAttrList = append(stgAttrList, stgAttr)
	}
	return stgAttrList, nil
}
