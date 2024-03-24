package main

import (
	"fmt"
	"os"

	"github.com/klaytn/klaytn/storage/database"
)

type DBEntryType uint8

const (
	MiscDB DBEntryType = iota
	headerDB
	BodyDB
	ReceiptsDB
	StateTrieDB
	StateTrieMigrationDB
	TxLookUpEntryDB
	bridgeServiceDB
	SnapshotDB
	// databaseEntryTypeSize should be the last item in this list!!
	databaseEntryTypeSize
)

func getDirs(path string) []string {
	dir, err := os.ReadDir(path) 
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	// stirng array
	results := make([]string, 0)

	for _, entry := range dir {
		if !entry.IsDir() {
			continue
		}
		results = append(results, entry.Name())
	}
	return results
}
func getLDB(dbc *database.DBConfig, dbEntryType uint8) database.Database {
	// if dbEntryType == uint8(StateTrieDB) {
	// 	dbm := database.NewDBManager(dbc)
	// 	return dbm.GetStateTrieDB()
	// }
	ldb, _ := database.NewLevelDB(dbc, database.DBEntryType(dbEntryType))
	return ldb
}
func getDBEntryType(entryName string) uint8 {
	if entryName == "misc" {
		return 0
	} else if entryName == "header" {
		return 1
	} else if entryName == "body" {
		return 2
	} else if entryName == "receipts" {
		return 3
	} else if entryName == "statetrie" {
		return 4
	} else if entryName == "statetrie_migrated" {
		return 5
	} else if entryName == "txlookup" {
		return 6
	} else if entryName == "bridgeservice" {
		return 7
	} else if entryName == "snapshot" {
		return 8
	} else {
		return 9
	}
}


func readDB(path string, dbEntryType uint8) {
	dbc := &database.DBConfig{
		Dir:                 path,
		DBType:              database.LevelDB,
		SingleDB:            false,
		NumStateTrieShards:  1,
		ParallelDBWrite:     false,
		OpenFilesLimit:      0,
		EnableDBPerfMetrics: false,
		LevelDBCacheSize:    0,
		RocksDBConfig:       nil,
		DynamoDBConfig:      nil,
	}
	ldb := getLDB(dbc, dbEntryType)
	defer ldb.Close()

	iter := ldb.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("[%x, %x]\n", key, value)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	dirs := getDirs("./db")

	for _, dir := range dirs {
		dbEntries := getDirs("./db/" + dir + "/output/cn1/data/klay/chaindata")
		fmt.Println("---------------------------------------------------")
		fmt.Println("> Found directories: " + dir)
		for _, dbEntry := range dbEntries {
			fmt.Println(">>>> Found DB Entry: " + dbEntry)
			dbEntryPath := "./db/" + dir + "/output/cn1/data/klay/chaindata/" + dbEntry
			readDB(dbEntryPath, getDBEntryType(dbEntry))
		}
	}
}
