package models

import (
	"fmt"
	"food_web/utils"
	"github.com/astaxie/beego"
	"strconv"
)
//用户收藏食品
func CollectFood(food Collect) (int64, error) {
	i, err := insertFood(food)
	return i, err
}

//储存收藏的食品id和用户名
func insertFood(food Collect) (int64, error) {
	return utils.ModifyDB("insert into user_shop_car(food_ids,user_ids) values(?,?)",
		food.Foodid,food.Userid)
}
//-----------查询文章---------

//根据页码查询文章
func FindFoodWithPage(page int) ([]Food_info, error) {
	//从配置文件中获取每页的文章数量
	num, _ := beego.AppConfig.Int("foodListPageNum")
	page--
	fmt.Println("---------->page", page)
	return QueryFoodWithPage(page, num)
}
/**
分页查询数据库
limit分页查询语句，
    语法：limit m，n

    m代表从多少位开始获取，与id值无关
    n代表获取多少条数据

注意limit前面没有where
*/
func QueryFoodWithPage(page, num int) ([]Food_info, error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return QueryFoodWithCon(sql)
}

func QueryFoodWithCon(sql string) ([]Food_info, error) {
	sql = "select id,food_store,food_pic,food_intro,food_price from food_info " + sql
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
//------翻页------

//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var foodRowsNum = 0

//只有首次获取行数的时候采取统计表里的行数
func GetFoodRowsNum() int {
	if foodRowsNum == 0 {
		foodRowsNum = QueryFoodRowNum()
	}
	return foodRowsNum
}

//查询文章的总条数
func QueryFoodRowNum() int {
	row := utils.QueryRowDB("select count(id) from food_info")
	num := 0
	row.Scan(&num)
	return num
}

//设置页数
func SetFoodRowsNum(){
	foodRowsNum = QueryFoodRowNum()
}

//根据id查询食品信息
func QueryFoodById(foodid string) []Food_info{
	sql :=fmt.Sprintf("select id,food_store,food_pic,food_intro,food_price from food_info where id='%s'",foodid)
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil
	}
	var foodList []Food_info
	for rows.Next() {
		id := foodid
		store := ""
		pic := ""
		intro := ""
		price := ""
		rows.Scan(&id, &store, &pic, &intro, &price)
		Id, _ := strconv.Atoi(id)
		food := Food_info{Id, store, pic, intro, price}
		foodList = append(foodList, food)
	}

	return foodList
}
//按条件查询用户,返回食品的id
func QueryFoodWithStore(con string) int {
	sql := fmt.Sprintf("select id from food_info %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	fmt.Println("food的id:",id)
	return id
}
//查询商品的评论
//根据id查询食品信息
func QueryFoodCommentById(food_id string) []Comment{
	sql :=fmt.Sprintf("select id,username,intro from comment where foodid='%s'",food_id)
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil
	}
	var commentList []Comment
	for rows.Next() {
		id := 0
		username := ""
		intro := ""
		foodid := food_id
		rows.Scan(&id, &username, &intro)
		comment := Comment{Username:username,Intro:intro,FoodId:foodid}
		commentList = append(commentList, comment)
	}

	return commentList
}