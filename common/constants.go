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
	PATH_MAIN_IMAGES string = "D:\\Projects\\gows\\src\\github.com\\phanhoc\\clonedb\\datatest\\images"
	SERVER                  = "https://www.sunfrog.com/search/?cId=0&cName=&search="
)

const MYSQL_CONN = "root:@tcp(localhost:3306)/sunfrog"
