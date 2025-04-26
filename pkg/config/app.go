package config
import(
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)
var(
	db *gorm.DB
)
func connect(){
	// dsn := "root:your_password@tcp(127.0.0.1:3306)/mysql"
	d,err:=gorm.Open("mysql","root:2004/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
	db=d
}
func GetDB() *gorm.DB{
	return db
}
