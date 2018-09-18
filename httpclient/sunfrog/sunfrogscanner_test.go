package sunfrog

import (
	"github.com/phanhoc/clonedb/common"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"
	"io/ioutil"
)

func TestSunfrog_GetData(t *testing.T) {
	scanner, _ := NewSunfrogScanner()
	res, err := scanner.GetData("https://www.sunfrog.com/search/paged4.cfm?cid=0&schSort=0&productType=0&search=shirts&offset=24")
	if err != nil {
		t.Fatalf("failed to get data, err %v", err)
	}
	t.Log(res)

}

func TestSunfrog_GetContinueUrl(t *testing.T) {
	data, err := ioutil.ReadFile("../../datatest/shirt.txt")
	if err != nil {
		t.Errorf("failed to read data from file, err: %v", err)
	}
	scanner, _ := NewSunfrogScanner()
	res, err := scanner.GetContinueUrl(string(data))
	if err != nil {
		t.Errorf("failed to get url from file, err: %v", err)
	}
	t.Log(res)
}

func TestTime(t *testing.T) {

	currentTime := time.Now().Local()
	subFolder := currentTime.Format("20060102")
	filename := path.Join(common.PATH_MAIN_IMAGES, subFolder, "shirt", "test.txt")
	path := filepath.Dir(filename)
	t.Log(path)
	if err := common.WriteDataToFile(filename, []byte("test")); err != nil {
		t.Logf("WriteDataToFile %v", err)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Log("here")
		if e := os.Mkdir(path, 0600); e != nil {
			t.Logf("err %v", e)
		}

	}
	_, err := os.Create(filename)
	if err != nil {
		t.Errorf("failed to create file: %s, err: %v", filename, err)
	}
}
