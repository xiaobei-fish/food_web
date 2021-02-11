package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/axgle/mahonia"
	"io"
	"net/http"
)
var (
	db orm.Ormer
)
type Food_info struct {
	Id 	  	  int
	FoodStore string
	FoodPic   string
	FoodIntro string
	FoodPrice string
}

func init(){
	orm.Debug = true //是否开启调试模式，调试模式下会打印SQL语句
	_ = orm.RegisterDataBase("default", "mysql", "root:qwe123@tcp(localhost:3306)/west?charset=utf8")
	orm.RegisterModel(new(Food_info))
	db = orm.NewOrm()
}
//食品信息插入数据库
func AddFood(food_Info *Food_info)(int64,error){
	id,err := db.Insert(food_Info)
	return id,err
}
//因为要多次检查错误，所以干脆自己建立一个函数。
func Check(err error){
	if err!=nil{
		fmt.Println(err)
	}
}
//获取源码
func HttpGet(url string) (result string, err error){
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println("HttpGet err:",err1)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	//爬取每页的链接
	for m:= 0 ; m < 40 ; m++{
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}

	return
}
//转化编码
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)

	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(tagCode)

	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result
}

