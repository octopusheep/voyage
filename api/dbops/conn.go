package dbops

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)


func init(){
	dbConn, err=sql.Open("mysql","root:Zhangyuyang1995@tcp(localhost:3306)/video_server")
	if err!=nil{
		panic(err.Error())
	}
}