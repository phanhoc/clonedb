package common

type Vendor int64

const (
	SUNFROG Vendor = iota
	AMAZONE
)

const (
	DATABASE_TYPE = "mysql"
	DATABASE_CONN = "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
)

const (
	PATH_MAIN_IMAGES    string = "./images"
	SUNFROG_SERVER             = "https://www.sunfrog.com/search/?cId=0&cName=&search="
	SUNFROG_SEARCH             = "https://www.sunfrog.com/search/"
	SUNFROG_BASE_PAGING int    = 24
)

const MYSQL_CONN = "root:@tcp(localhost:3307)/sunfrog"
