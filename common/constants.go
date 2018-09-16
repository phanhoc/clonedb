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
	PATH_MAIN_IMAGES string = "./datatest/images"
	SERVER                  = "https://www.sunfrog.com/search/?cId=0&cName=&search=shirt"
)
