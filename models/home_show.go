package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)

type HomeBlockParam struct {
	Id         int
	Store      string
	Intro      string
	Pic    	   string
	Price      string

	//查看文章的地址
	Link string

	//记录是否登录
	IsLogin bool
}
//分页的结构体
type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

//----------首页显示内容---------
func MakeHomeBlocks(foods []Food_info, isLogin bool) template.HTML {
	htmlHome := ""
	for _, food := range foods {

		//将数据库model转换为首页模板所需要的model
		homeParam 		:= HomeBlockParam{}
		homeParam.Id 	= food.Id
		homeParam.Store = food.FoodStore
		homeParam.Intro = food.FoodIntro
		homeParam.Pic 	= food.FoodPic
		homeParam.Price = food.FoodPrice

		homeParam.Link = "/food/" + strconv.Itoa(food.Id)

		homeParam.IsLogin = isLogin

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}
//----------购物车显示内容---------
func MakeCarBlocks(foods []Food_info, isLogin bool) template.HTML {
	htmlHome := ""
	for _, food := range foods {

		//将数据库model转换为首页模板所需要的model
		homeParam 		:= HomeBlockParam{}
		homeParam.Id 	= food.Id
		homeParam.Store = food.FoodStore
		homeParam.Intro = food.FoodIntro
		homeParam.Pic 	= food.FoodPic
		homeParam.Price = food.FoodPrice

		homeParam.Link = "/food/" + strconv.Itoa(food.Id)

		homeParam.IsLogin = isLogin

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/car_block.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}

//-----------翻页-----------
//page是当前的页数
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}

	//查询出总的条数
	num := GetFoodRowsNum()

	//从配置文件中读取每页显示的条数
	pageRow, _ := beego.AppConfig.Int("foodListPageNum")

	//计算出总页数
	fmt.Println(num)
	allPageNum := (num-1)/pageRow + 1

	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//当前页数小于等于1，那么上一页的按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于等于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}


