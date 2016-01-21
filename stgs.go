// goStorer
package main

import (
	"log"
)

const (
	GOSTORER_HOME = ""
)

func Store(stgUuid string, doc string) (bool, string) {
	var (
		ok       bool
		errormsg string
	)
	ps := getPStorage(stgUuid)
	if ps != nil {
		e := (*ps).ExternalStorage_Store(doc)
		if e != nil {
			log.Println(e.Error())
		}
	} else {
		ok, errormsg = false, "Unknown Storage (StorageUuid = "+stgUuid+")"
		log.Println(errormsg)
	}
	return ok, errormsg
}

func main() {
	log.Println("Starting stgs.")

	err := constructStorageMap(GOSTORER_HOME + "config/storages.xml")
	if err != nil {
		log.Println("Failed to construct map of storages: " + err.Error())
		return
	}

	log.Println("Successfully started goStorer")
	Store("uuid001", "Junimond")
	Store("uuid002", "Julimond")
	Store("uuid003", "Wintermärchen")
	Store("uuid004", "xxx")
	Store("uuid001", "Junimond")

	log.Println("stgs stopped.")
}
