# 生成api文档
goctl api go -api greet.api -dir . -style gozero



#导入gorm
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
#run
go run greet.go -f etc/greet-api.yaml


#表结构
type LoginRequest {
Name     string `json:"name"`
Password string `json:"password"`
}
type UserRegisterRequest {
Name string `json:"name"`
Age  int    `json:"age"`
}
goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/gormtest" -table="*"  -dir="./model"

