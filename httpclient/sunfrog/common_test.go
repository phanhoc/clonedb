package sunfrog

import (
	"io/ioutil"
	"testing"
)

func TestSunfrog_GetTitleNiche(t *testing.T) {
	data, err := ioutil.ReadFile("../../datatest/one_shirt.txt")
	if err != nil {
		t.Fatalf("failed to read data from file, err: %v", err)
	}
	res, err := getTitleNiche(string(data))
	if err != nil {
		t.Fatalf("failed to get title from data, err: %v", err)
	}
	t.Log(res)
}

func TestSunfrog_GetDescriptionNiche(t *testing.T) {
	data, err := ioutil.ReadFile("../../datatest/one_shirt.txt")
	if err != nil {
		t.Fatalf("failed to read data from file, err: %v", err)
	}
	res, err := getDescriptionNiche(string(data))
	if err != nil {
		t.Fatalf("failed to get description from data, err: %v", err)
	}
	t.Log(res)
}

func TestSunfrog_GetUrlNiche(t *testing.T) {
	data, err := ioutil.ReadFile("../../datatest/one_shirt.txt")
	if err != nil {
		t.Fatalf("failed to read data from file, err: %v", err)
	}
	res, err := getUrlNiche(string(data))
	if err != nil {
		t.Fatalf("failed to get url from data, err: %v", err)
	}
	t.Log(res)
}

func TestSunfrog_GetImageNiche(t *testing.T) {
	data, err := ioutil.ReadFile("../../datatest/one_shirt.txt")
	if err != nil {
		t.Fatalf("failed to read data from file, err: %v", err)
	}
	res, err := getMainImageNiche(string(data), "shirt")
	if err != nil {
		t.Fatalf("failed to get url from data, err: %v", err)
	}
	t.Log(res)
}

func TestSunfrog_GetMoneyNiche(t *testing.T) {
	data, err := ioutil.ReadFile("../../datatest/one_normal_shirt.txt")
	if err != nil {
		t.Fatalf("failed to read data from file, err: %v", err)
	}
	res, err := getMoneyNiche(string(data))
	if err != nil {
		t.Fatalf("failed to get url from data, err: %v", err)
	}
	t.Log(res)
}

func TestSunfrog_GetContentNiche(t *testing.T) {
	data, err := ioutil.ReadFile("../../datatest/one_shirt.txt")
	if err != nil {
		t.Fatalf("failed to read data from file, err: %v", err)
	}
	res, err := getContentNiche(string(data))
	if err != nil {
		t.Fatalf("failed to get url from data, err: %v", err)
	}
	t.Log(res)
}
