package sunfrog

import (
	"fmt"
	"github.com/ddliu/go-httpclient"
	"github.com/phanhoc/clonedb/model/sun"
	"regexp"
	"time"
)

const (
	USERAGENT       = "Mozilla/5.0 (X11; U; Linux i686; en-US) AppleWebKit/532.4 (KHTML, like Gecko) Chrome/4.0.233.0 Safari/532.4"
	TIMEOUT         = 300
	CONNECT_TIMEOUT = 5
	REF             = "http://www.google.com/"
)

type Sunfrog struct {
}

func NewSunfrogScanner() (*Sunfrog, error) {
	return &Sunfrog{}, nil
}

func (s *Sunfrog) GetData(url string) (string, error) {
	httpclient.Defaults(httpclient.Map{
		"opt_useragent":   USERAGENT,
		"opt_timeout":     TIMEOUT,
		"Accept-Encoding": "gzip,deflate,sdch",
	})

	res, err := httpclient.
		WithHeader("Accept-Language", "en-us").
		WithHeader("Referer", REF).
		Get(url)

	if err != nil {
		return "", fmt.Errorf("failed to get data, err: %v", err)
	}

	//fmt.Println("Response:")
	body, err := res.ReadAll()
	if err != nil {
		return "", fmt.Errorf("failed to read body data, err: %v", err)
	}
	//ioutil.WriteFile("./datatest/one_shirt.txt", []byte(string(body)), 0600)

	return string(body), nil

}

func (s *Sunfrog) GetAllUrl(data string) ([]string, error) {
	//data, err := ioutil.ReadFile("./datatest/shirt.txt")
	//if err != nil {
	//	return nil, fmt.Errorf("failed to read data from file, err: %v", err)
	//}
	regexUrl := "<a href=\"(.*?)\"  border=\"0\""
	re, _ := regexp.Compile(regexUrl)
	result := re.FindAllStringSubmatch(data, -3)
	res := make([]string, 0, len(result))
	for _, item := range result {
		res = append(res, item[1])
	}

	return res, nil
}

func (s *Sunfrog) GetDetailNiche(data string) (interface{}, error) {
	sunTShirt := new(sun.TShirt)
	url, err := getUrlNiche(data)
	if err != nil {
		return nil, err
	}
	title, err := getTitleNiche(data)
	if err != nil {
		return nil, err
	}
	image, err := getMainImageNiche(data)
	if err != nil {
		return nil, err
	}
	description, err := getDescriptionNiche(data)
	if err != nil {
		return nil, err
	}
	money, err := getMoneyNiche(data)
	if err != nil {
		return nil, err
	}
	content, err := getContentNiche(data)
	if err != nil {
		return nil, err
	}
	sunTShirt.Url = url
	sunTShirt.Title = title
	sunTShirt.Images = image
	sunTShirt.Desc = description
	sunTShirt.Money = money
	sunTShirt.Content = content
	sunTShirt.Time = time.Now().String()

	return sunTShirt, nil
}
