/*
* Http (curl) request in golang
* @author phan hoc
*/
package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/phanhoc/clonedb/common"
	sc "github.com/phanhoc/clonedb/httpclient"
	"github.com/phanhoc/clonedb/model/sun"
)

// url := "https://www.sunfrog.com/search/?cId=0&cName=&search=shirt"

func main() {
	fmt.Println(fmt.Sprintf("Starting clone data from %s", common.SERVER))
	c := make(chan []*sun.TShirt)
	go pullData(c)

	listNiche := <-c
	fmt.Println(spew.Sdump(listNiche))
}

func pullData(req chan []*sun.TShirt) {
	listNiche, err := requestData()
	if err != nil {
		fmt.Println("failed to request data from sunfrog, err: %v", err)
	} else {
		req <- listNiche
	}
}

func requestData() ([]*sun.TShirt, error) {
	scanner, err := sc.NewScanner(common.SUNFROG)
	if err != nil {
		return nil, fmt.Errorf("failed to create new scanner, err: %v", err)
	}
	res, err := scanner.GetData(common.SERVER)
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
		detail, err := scanner.GetDetailNiche(dataNiche)
		if err != nil {
			return nil, fmt.Errorf("failed to parse detail niche, err: %v", err)
		}
		detailNiche := detail.(*sun.TShirt)
		listNiche = append(listNiche, detailNiche)
	}

	return listNiche, nil
}
