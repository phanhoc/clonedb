package sunfrog

import (
	"fmt"
	"github.com/phanhoc/clonedb/common"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
	"errors"
)

func getTitleNiche(data string) (string, error) {
	titleRegex := "<meta property=\"og:title\" content=\"(.*?)\"/>"

	return common.ParseData(titleRegex, data)
}

func getDescriptionNiche(data string) (string, error) {
	descriptionRegex := "<meta property=\"og:description\" content=\"(.*?)\"/>"

	return common.ParseData(descriptionRegex, data)
}

func getUrlNiche(data string) (string, error) {
	urlRegex := "<meta property=\"og:url\" content=\"(.*?)\"/>"

	return common.ParseData(urlRegex, data)
}

func getMainImageNiche(data, key string) (string, error) {
	mainImageRegex := "<meta property=\"og:image\" content='(.*?)'/>"
	result, err := common.ParseData(mainImageRegex, data)
	if err != nil {
		return "", fmt.Errorf("failed to parse data, regex: %s, err: %v", mainImageRegex, err)
	}

	response, err := http.Get(result)
	if err != nil {
		return "", fmt.Errorf("failed to curl to url: %s, err: %v", result, err)
	}

	defer response.Body.Close()
	name := path.Base(result)
	currentTime := time.Now().Local()
	subFolder := currentTime.Format("20060102")
	filename := path.Join(common.PATH_MAIN_IMAGES, subFolder, key, name)
	path := filepath.Dir(filename)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %s, err: %v", filename, err)
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %s, err: %v", filename, err)
	}
	file.Close()

	return filename, nil
}

func getMoneyNiche(data string) (string, error) {
	moneyRegex := "<span style=\"text-decoration: line-through; margin-right:4px;\""
	afterRegex := "</strong>"
	firstIndex := strings.Index(data, moneyRegex)
	if firstIndex == -1 {
		moneyRegex = "<strong style=\"display:block; margin-bottom:0px; "
		firstIndex = strings.Index(data, moneyRegex)
		if firstIndex == -1 {
			return "", errors.New("unable to find sub string")
		}
		afterData := data[firstIndex+len(moneyRegex):]
		lastIndex := strings.Index(afterData, afterRegex)
		moneyData := afterData[:lastIndex]
		moneyIndex := strings.LastIndex(moneyData, "$")
		return strings.TrimSpace(moneyData[moneyIndex+1:]), nil
	} else {
		afterData := data[firstIndex+len(moneyRegex):]
		lastIndex := strings.Index(afterData, afterRegex)
		moneyData := afterData[:lastIndex]
		moneyIndex := strings.LastIndex(moneyData, "$")
		return strings.TrimSpace(moneyData[moneyIndex+1:]), nil
	}

}

func getContentNiche(data string) (string, error) {
	return getDescriptionNiche(data)
}
