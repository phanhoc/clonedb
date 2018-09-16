package sunfrog

import (
	"errors"
	"fmt"
	"github.com/phanhoc/clonedb/common"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

func parseData(regex, data string) (string, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return "", fmt.Errorf("failed to create regex expresion, regex: %s, err: %v", regex, err)
	}
	result := re.FindStringSubmatch(data)
	if len(result) < 2 {
		return "", fmt.Errorf("unable to find sub string")
	}
	return result[1], nil
}

func getTitleNiche(data string) (string, error) {
	titleRegex := "<meta property=\"og:title\" content=\"(.*?)\"/>"

	return parseData(titleRegex, data)
}

func getDescriptionNiche(data string) (string, error) {
	descriptionRegex := "<meta property=\"og:description\" content=\"(.*?)\"/>"

	return parseData(descriptionRegex, data)
}

func getUrlNiche(data string) (string, error) {
	urlRegex := "<meta property=\"og:url\" content=\"(.*?)\"/>"

	return parseData(urlRegex, data)
}

func getMainImageNiche(data string) (string, error) {
	mainImageRegex := "<meta property=\"og:image\" content='(.*?)'/>"
	result, err := parseData(mainImageRegex, data)
	if err != nil {
		return "", fmt.Errorf("failed to parse data, regex: %s, err: %v", mainImageRegex, err)
	}

	response, err := http.Get(result)
	if err != nil {
		return "", fmt.Errorf("failed to curl to url: %s, err: %v", result, err)
	}

	defer response.Body.Close()
	name := path.Base(result)
	filename := path.Join(common.PATH_MAIN_IMAGES, name)
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
		return "", errors.New("unable to find sub string")
	}
	afterData := data[firstIndex+len(moneyRegex):]
	lastIndex := strings.Index(afterData, afterRegex)
	moneyData := afterData[:lastIndex]
	moneyIndex := strings.LastIndex(moneyData, "$")
	return strings.TrimSpace(moneyData[moneyIndex+1:]), nil
}

func getContentNiche(data string) (string, error) {
	return "", nil
}
