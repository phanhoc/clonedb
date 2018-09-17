package common

type Vendor int64

const (
	SUNFROG Vendor = iota
	AMAZONE
)

const (
	PATH_MAIN_IMAGES string = "./datatest/images"
	SERVER                  = "https://www.sunfrog.com/search/?cId=0&cName=&search="
)

const MYSQL_CONN = "root:@tcp(localhost:3307)/sunfrog"
