package httpclient

import (
	"errors"
	"github.com/phanhoc/clonedb/common"
	"github.com/phanhoc/clonedb/httpclient/sunfrog"
)

type Scanner interface {
	GetData(string) (string, error)
	GetAllUrl(string) ([]string, error)
	GetDetailNiche(string, string) (interface{}, error)
}

func NewScanner(vendor common.Vendor) (Scanner, error) {
	switch vendor {
	case common.SUNFROG:
		return sunfrog.NewSunfrogScanner()
	default:
		return nil, errors.New("partner unsuported")
	}

}
