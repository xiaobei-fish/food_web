package models

import (
	"fmt"
	"food_web/utils"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type User struct {
	Id		   string
	Username   string
	Password   string
	Salt	   string
	Genre	   string // 0 买家， 1 卖家， 2 管理员
	Status     int 	  // 0 正常状态， 1删除
	Createtime int64
}
type Collect struct {
	Userid string
	Foodid string
}
type Seller struct {
	Username string
	FoodId	 string
}
type UserMessage struct{
	Id		int
	UserId	int
	Words	string
	Picture string
	Sex		string
}
//插入新用户数据到数据库
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,salt,genre,status,createtime) values (?,?,?,?,?,?)",
		user.Username, user.Password, user.Salt, user.Genre, user.Status, user.Createtime)
}
//插入卖家信息到数据库
func InsertSeller(seller Seller) (int64, error) {
	return utils.ModifyDB("insert into sellers(username,foodid) values(?,?)",
		seller.Username,seller.FoodId)
}
//按条件查询用户,返回用户的id
func QueryUserWithCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	fmt.Println("user的id:",id)
	return id
}

//查询用户的类型
func QueryUserGenre(username string) bool {
	sql := fmt.Sprintf("select id from users where username='%s' and genre='1'",username)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	if id > 0 {
		return true
	}else{
		return false
	}
}

//按条件查询用户购物车
func QueryUserCollectWithCon(con string) int {
	sql := fmt.Sprintf("select id from user_shop_car %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWithCon(sql)
}

//根据用户种类，查询id
func QueryUserWithGenre(username, password, genre string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s' and (genre = '%s')", username, password, genre)
	return QueryUserWithCon(sql)
}

//根据用户id，检查是否已经收藏
func QueryFoodWithUserId(userId, foodId string) int {
	sql := fmt.Sprintf("where user_ids='%s' and food_ids='%s'", userId, foodId)
	return QueryUserCollectWithCon(sql)
}

var usercileRowsNum  = 0

//只有首次获取行数的时候采取统计表里的行数,好像没啥用,前端处理好了
/*func GetUserRowsNum() int {
	if usercileRowsNum == 0 {
		usercileRowsNum = QueryNovelRowNum()
	}
	return usercileRowsNum
}*/

//设置用户页数
func SetUserRowsNum(){
	usercileRowsNum = QueryUserRowNum()
}

//查询用户的条数
func QueryUserRowNum() int {
	row := utils.QueryRowDB("select count(id) from users")
	num := 0
	row.Scan(&num)
	return num
}
//查询用户购物车
func FindCarWithUserId(userid string) ([]Food_info, error) {
	Id := QueryUserWithUsername(userid)
	condition := "(select food_ids from user_shop_car where user_ids="
	userId := strconv.Itoa(Id) + ")"
	sql := "select id,food_store,food_pic,food_intro,food_price from food_info where id in " + condition + userId
	//fmt.Println(sql)
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var foodList []Food_info
	for rows.Next() {
		id := 0
		store := ""
		pic := ""
		intro := ""
		price := ""
		rows.Scan(&id, &store, &pic, &intro, &price)
		food := Food_info{id, store, pic, intro, price}
		foodList = append(foodList, food)
	}
	return foodList, nil
}
//删除购物车内的食品
func DeleteCarFoodWithId(foodID string,userID string) (int64, error) {
	sql:= "delete from user_shop_car where food_ids=" + foodID + " and user_ids=" + userID
	fmt.Println(sql)
	return utils.ModifyDB(sql)
}
/*
//根据页码查询用户并返回数据以便展示
func FindUserWithPage(page,limit int) []orm.Params {
	orm.Debug = true

	page--
	sql := fmt.Sprintf("limit %d,%d", page*limit,limit)
	//fmt.Println("sql:::::::::::",sql)
	sql = "select id,username,genre from users " + sql
	//fmt.Println("sql:::::::::::",sql)

	var Users []orm.Params
	fmt.Println("var是正常的")
	i, e := db.Raw(sql).Values(&Users)
	fmt.Println("输出是:",i)
	if e != nil {
		fmt.Println("Raw出错")
	}

	return Users
}*/
//个人页面信息操作
//个人信息插入数据库
func AddMes(user_message UserMessage)(int64,error){
	return utils.ModifyDB("insert into user_message(user_id,picture,sex,words) values(?,?,?,?)",
		user_message.UserId,user_message.Picture,user_message.Sex,user_message.Words)
}
func QueryUserMessageBool(user_id int) bool{
	sql := "select id from user_message where user_id='" + strconv.Itoa(user_id) + "'"
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	if id > 0 {
		return true
	}else{
		return false
	}
}
//查询用户的类型
func QueryGenre(username string) int {
	sql := fmt.Sprintf("select id from users where username='%s' and genre='1'",username)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	if id > 0 {
		return 1
	}else{
		sql := fmt.Sprintf("select id from users where username='%s' and genre='0'",username)
		fmt.Println(sql)
		row := utils.QueryRowDB(sql)
		id := 0
		row.Scan(&id)
		if id > 0 {
			return 0
		}else{
			sql := fmt.Sprintf("select id from users where username='%s' and genre='2'",username)
			fmt.Println(sql)
			row := utils.QueryRowDB(sql)
			id := 0
			row.Scan(&id)
			if id > 0 {
				return 2
			}else{
				return -1
			}
		}
	}
}
//根据id查询用户信息
func QueryUserMessageById(user_id string) []UserMessage{
	sql :=fmt.Sprintf("select id,picture,words,sex from user_message where user_id='%s'",user_id)
	fmt.Println(sql)
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil
	}
	var uList []UserMessage
	for rows.Next() {
		id := 0
		userid := user_id
		words := ""
		pic := ""
		sex := ""
		rows.Scan(&id, &pic, &words, &sex)
		userId,_ := strconv.Atoi(userid)
		mes := UserMessage{Id:id,UserId:userId,Picture:pic,Words:words,Sex:sex}
		uList = append(uList, mes)
	}
	return uList
}

//修改密码
func AlterPassword(username,password string) (int64, error){
	orm.Debug = true
	password = utils.MD5(password)
	salt := utils.SaltMD5(password)
	sql := "update users SET password='" + password + "', salt='" + salt + "' "
	location := "where username='" + username + "'"
	sql = sql + location
	fmt.Println("sql::::",sql)
	return utils.ModifyDB(sql)
}

//验证旧密码是否输入一致
func TestOldPassword(oldPassword string) int {
	orm.Debug = true
	oldPassword = utils.MD5(oldPassword)
	sql := fmt.Sprintf("select id from users where password='%s'", oldPassword)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	fmt.Println("user的id:", id)
	return id
}