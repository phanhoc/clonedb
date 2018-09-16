package sunfrog

import "testing"

func TestSunfrog_GetData(t *testing.T) {
	scanner, _ := NewSunfrogScanner()
	res, err := scanner.GetData("https://www.sunfrog.com/search/paged4.cfm?cid=0&schSort=0&productType=0&search=dogs&offset=1204")
	if err != nil {
		t.Fatalf("failed to get data, err %v", err)
	}
	t.Log(res)

}
