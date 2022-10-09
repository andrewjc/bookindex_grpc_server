package indexer

import (
	"log"
	"os"
	"time"
)

type Indexer struct {
	config *map[string]interface{}
}

func IndexProcess(config *map[string]interface{}) *Indexer {
	return &Indexer{config}
}

func (indexer *Indexer) Start() {
	indexer.PerformScan(time.Now().Local())
	for theTime := range time.Tick(time.Minute * 15) {
		indexer.PerformScan(theTime)
	}
}

func (indexer *Indexer) PerformScan(startTime time.Time) {
	dirName := (*indexer.config)["folder"].(string)
	log.Printf("Performing scan of %s at %v", dirName, startTime)

	entries, err := os.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		indexer.processFile(dirName, entry.Name())
	}

}

func (indexer *Indexer) processFile(path string, filename string) {
	log.Println(path, filename)
}
