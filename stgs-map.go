// map of external storages (safe)
package main

import (
	"github.com/s-pet/storages/storage"
	"log"
	"sync"
)

type StorageMap struct {
	sync.RWMutex
	byUuid map[string]*storage.ExternalStorage
}

var (
	storages = struct {
		sync.RWMutex
		byUuid map[string]*storage.ExternalStorage
	}{byUuid: make(map[string]*storage.ExternalStorage)}
)

func constructStorageMap(xmlFilename string) error {
	stgAttrList, err := readStorageConfig(xmlFilename)
	if stgAttrList != nil {
		for _, stgAttr := range stgAttrList {
			stg, err := storage.InitExternalStorage(stgAttr)
			if err == nil {
				addStorage(stg)
			}
		}
	}
	return err
}

func addStorage(s storage.ExternalStorage) {
	if s != nil {
		storages.Lock()
		storages.byUuid[s.ExternalStorage_GetUuid()] = &s
		log.Print("ADDED ", s.ExternalStorage_GetUuid())
		storages.Unlock()
	}
}

func getPStorage(stgUuid string) *storage.ExternalStorage {
	storages.RLock()
	ps := storages.byUuid[stgUuid]
	storages.RUnlock()
	return ps
}
