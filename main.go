/*
* Http (curl) request in golang
* @author phan hoc
 */
package main

import (
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/phanhoc/clonedb/common"
	"github.com/phanhoc/clonedb/db"
	"github.com/phanhoc/clonedb/db/gorm"
	sc "github.com/phanhoc/clonedb/httpclient"
	"github.com/phanhoc/clonedb/model/sun"
	"log"
)

// url := "https://www.sunfrog.com/search/?cId=0&cName=&search=shirt"

func main() {
	key := "shirt"
	url := common.SERVER + key
	fmt.Println(fmt.Sprintf("Starting clone data from %s", key))
	sqlDB, err := setupDb()
	if err != nil {
		log.Fatalf("failed to setup db, %v", err)
	}
	c := make(chan []*sun.TShirt)
	go pullData(url, key, c)

	listNiche := <-c
	fmt.Println(spew.Sdump(listNiche))
	if err := saveData(sqlDB, listNiche); err != nil {
		log.Fatalf("failed to save db, %v", err)
	}
}

func pullData(url, key string, req chan []*sun.TShirt) {
	listNiche, err := requestData(url, key)
	if err != nil {
		log.Fatalf("failed to request data from sunfrog, err: %v", err)
	} else {
		req <- listNiche
	}
}

func saveData(db db.DB, listNiche []*sun.TShirt) error {
	if len(listNiche) < 1 {
		return errors.New("invalidate agrement")
	}

	for _, niche := range listNiche {
		err := db.InsertNiche(niche)
		if err != nil {
			return fmt.Errorf("failed to save niche: %s, err: %v", niche.Title, err)
		}
	}
	return nil
}

func setupDb() (db.DB, error) {
	sqlDb, err := gorm.NewDB("mysql", common.MYSQL_CONN)
	if err != nil {
		return nil, err
	}
	err = sqlDb.MigrateSchema()
	if err != nil {
		return nil, err
	}
	return sqlDb, nil
}

func requestData(url, key string) ([]*sun.TShirt, error) {
	scanner, err := sc.NewScanner(common.SUNFROG)
	if err != nil {
		return nil, fmt.Errorf("failed to create new scanner, err: %v", err)
	}
	res, err := scanner.GetData(url)
	if err != nil {
		return nil, fmt.Errorf("failed to search all niche by key, err: %v", err)
	}
	listUrl, err := scanner.GetAllUrl(res)
	if err != nil {
		return nil, fmt.Errorf("failed to list all url, err: %v", err)
	}
	listNiche := make([]*sun.TShirt, 0, len(listUrl))
	for _, url := range listUrl {
		fmt.Println(fmt.Sprintf("query detail niche from url: %s", url))
		dataNiche, err := scanner.GetData(url)
		if err != nil {
			return nil, fmt.Errorf("failed to query detail niche, url: %s, err: %v", url, err)
		}
		detail, err := scanner.GetDetailNiche(dataNiche, key)
		if err != nil {
			return nil, fmt.Errorf("failed to parse detail niche, err: %v", err)
		}
		detailNiche := detail.(*sun.TShirt)
		listNiche = append(listNiche, detailNiche)
	}

	return listNiche, nil
}
