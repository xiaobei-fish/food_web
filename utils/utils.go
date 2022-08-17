package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/logs"
	"log"
)

var db *sql.DB
var salt = "abrghikjmnxdwqopABRGHIKJMNXDWQOP"

func InitMysql() {
	fmt.Println("InitMysql....")
	driverName := "mysql"

	dbConn := "root" + ":" + "qwe123" + "@tcp(" + "localhost" + ":" + "3306" + ")/" + "west" + "?charset=utf8"
	fmt.Println(dbConn)
	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreateTableWithUser()
		CreateTableNamedByUserShopCar()
		CreteTableWithSeller()
		CreateTableWithUserMes()
		CreateTableNamedByUserBag()
		CreteTableWithMoney()
		CreteTableWithComment()
	}
}

//日志开启
func LogOn(){
	log := logs.NewLogger(10000) //创建一个日志记录器，参数为缓冲区的大小
	// 设置配置文件
	jsonConfig := `{
        "filename" : "../test.log", 
        "maxlines" : 1000,      
        "maxsize"  : 10240
    }`
	log.SetLogger("file", jsonConfig) 	//设置日志记录方式：本地文件记录
	log.SetLevel(logs.LevelDebug)     				//设置日志写入缓冲区的等级
	log.EnableFuncCallDepth(true)     			//输出log时能显示输出文件名和行号

	log.Emergency("Emergency")
	log.Alert("Alert")
	log.Critical("Critical")
	log.Error("Error")
	log.Warning("Warning")
	log.Notice("Notice")
	log.Informational("Informational")
	log.Debug("Debug")

	log.Flush() 									//将日志从缓冲区读出，写入到文件
	log.Close()
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	//fmt.Println("sql::",sql)
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		salt VARCHAR(64),
		genre INT(4),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

//创建个人页面表
func CreateTableWithUserMes() {
	sql := `CREATE TABLE IF NOT EXISTS user_message(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		user_id INT(4),
		picture VARCHAR(64),
		words	VARCHAR(64),
		sex		VARCHAR(64),
		CONSTRAINT fk_user_message_users FOREIGN KEY (user_id) REFERENCES users (id)  on delete restrict   ON UPDATE CASCADE
		)AUTO_INCREMENT = 1;`
	ModifyDB(sql)
}

//创建用户购物车表
func CreateTableNamedByUserShopCar()  {
	sql := `CREATE TABLE IF NOT EXISTS user_shop_car(
		id INT(11)  NOT NULL PRIMARY KEY AUTO_INCREMENT,
		food_ids INT(10) NOT NULL,
		user_ids INT(10),
		CONSTRAINT fk_user_shop_car_food_info FOREIGN KEY (food_ids) REFERENCES food_info (id)  on delete restrict   ON UPDATE CASCADE
		)AUTO_INCREMENT = 1;`
	ModifyDB(sql)
}

//创建用户背包表
func CreateTableNamedByUserBag()  {
	sql := `CREATE TABLE IF NOT EXISTS user_bag(
		id INT(11)  NOT NULL PRIMARY KEY AUTO_INCREMENT,
		food_store VARCHAR(100),
		food_pic VARCHAR(200),
		food_intro VARCHAR(200),
		food_price VARCHAR(50),
		user_ids INT(10)
		)AUTO_INCREMENT = 1;`
	ModifyDB(sql)
}

//创建卖家表
func CreteTableWithSeller(){
	sql := `CREATE TABLE IF NOT EXISTS sellers(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		foodid   VARCHAR(64)
		);`
	ModifyDB(sql)
}

//创建货币表
func CreteTableWithMoney(){
	sql := `CREATE TABLE IF NOT EXISTS money(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		lmoney   VARCHAR(64)
		);`
	ModifyDB(sql)
}

//创建评论表
func CreteTableWithComment(){
	sql := `CREATE TABLE IF NOT EXISTS comment(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		foodid   VARCHAR(64),
		intro	 VARCHAR(200)
		);`
	ModifyDB(sql)
}

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

//MD5加密密码
func MD5(str string) string{
	loca := len(str)
	str = str + salt[loca-1 : loca]
	md5str := fmt.Sprintf("%x",md5.Sum([]byte(str)))
	return  md5str
}

//随机盐MD5值
func SaltMD5(str string) string{
	loca := len(str)
	md5salt := salt[loca-1 : loca]
	md5str := fmt.Sprintf("%x",md5.Sum([]byte(md5salt)))
	return  md5str
}

