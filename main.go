/*
* Http (curl) request in golang
* @author phan hoc
 */
package main

import (
	"errors"
	"fmt"
	"github.com/phanhoc/clonedb/common"
	"github.com/phanhoc/clonedb/db"
	"github.com/phanhoc/clonedb/db/gorm"
	sc "github.com/phanhoc/clonedb/httpclient"
	"github.com/phanhoc/clonedb/model/sun"
	"github.com/jessevdk/go-flags"
	"os"
	"strings"
	"strconv"
	"nextop/c-horde/walletengine/core/app"
	"syscall"
)

var (
	newlineBytes = []byte{'\n'}
)

func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Stderr.Write(newlineBytes)
	os.Exit(1)
}

func logf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Stdout.Write(newlineBytes)
}

func main() {
	server := app.NewApp(
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	server.Run(func() {
		var opts = struct {
			Url      string `short:"u" long:"url" description:"Url to get data of niche"`
			Key      string `short:"k" long:"key" description:"Key to search niche"`
			Category string `short:"c" long:"category" description:"Category of niche"`
		}{
			Url: common.SUNFROG_SERVER,
		}
		_, err := flags.Parse(&opts)
		if err != nil {
			fatalf("%v", err)
		}
		if opts.Key == "" {
			fatalf("please enter key to searching niche. Using clonedb -h or clonedb --help to more information")
		}

		url := opts.Url + opts.Key
		if opts.Category != "" {
			if strings.Contains(url, "cName=") {
				strings.Replace(url, "cName=", strings.Join([]string{"cName="}, opts.Category), 1)
			}
		}
		fmt.Println(fmt.Sprintf("Starting clone data from %s", url))
		sqlDB, err := setupDb()
		if err != nil {
			fatalf("failed to setup db, %v", err)
		}
		c := make(chan []*sun.TShirt)
		go pullData(url, opts.Key, c)

		go handlerDB(c, sqlDB)

		server.AtExit(func() {
			logf("Stopping application")
		})
	})
	server.Wait()
	logf("Shutdown completed, uptime %s", server.Elapsed())

}

func pullData(url, key string, req chan []*sun.TShirt) {
	i := int(0)
	scanner, err := sc.NewScanner(common.SUNFROG)
	if err != nil {
		fatalf("failed to create new scanner, err: %v", err)
	}
	for {
		i++
		urlInternal := url
		res, err := scanner.GetData(urlInternal)
		if err != nil {
			logf("failed to search all niche by key, err: %v", err)
		}
		listNiche, err := requestData(scanner, res, key)
		if err != nil {
			logf("failed to request data from sunfrog, err: %v", err)
		} else {
			req <- listNiche
		}
		uc, err := scanner.GetContinueUrl(res)
		if err != nil {
			logf("failed to get url paging, err: %v", err)
		}
		urlInternal = uc + strconv.Itoa(i*common.SUNFROG_BASE_PAGING)
		fmt.Println("urlInternal ", urlInternal)
	}

}

func handlerDB(req chan []*sun.TShirt, db db.DB) {
	for {
		select {
		case listNiche := <-req:
			// fmt.Println(spew.Sdump(listNiche))
			if err := saveData(db, listNiche); ergitr != nil {
				logf("failed to save db, %v", err)
			}
		}
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

func requestData(scanner sc.Scanner, data, key string) ([]*sun.TShirt, error) {
	// scanner, err := sc.NewScanner(common.SUNFROG)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create new scanner, err: %v", err)
	// }
	// res, err := scanner.GetData(url)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to search all niche by key, err: %v", err)
	// }
	listUrl, err := scanner.GetAllUrl(data)
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
