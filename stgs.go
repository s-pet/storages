// goStorer
package main

import (
	"log"
	"os"
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
	log.Println("Starting storages...")

	//CONFIGURATION OF EXTERNAL STORAGES
	//configuration file: <current-working-directory>/config/storages.xml
	//read configuration file and construct map of external storages
	cwd, errCwd := os.Getwd()
	if errCwd != nil {
		log.Fatal("Cannot determine current working directory: " + errCwd.Error())
	}
	storages_config := cwd + "/config/storages.xml"
	finfo, errFinfo := os.Stat(storages_config)
	if errFinfo != nil {
		log.Fatal("Cannot access storage configuration: " + errFinfo.Error())
	}
	if finfo.IsDir() {
		log.Fatal("Cannot access storage configuration: " + errFinfo.Error())
	}
	err := constructStorageMap(storages_config)
	if err != nil {
		log.Fatal("Cannot construct map of storages: " + err.Error())
	}

	Store("UUID-DOES-NOT-EXIST", "")
	Store("DUMMY-UUID-NasExample", "dummy")

	log.Println("storages successfully stopped.")
}
