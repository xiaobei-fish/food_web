package controllers

import (
	_ "database/sql"
	"fmt"
	"food_web/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	_ "github.com/axgle/mahonia"
	_ "io"
	_ "net/http"
	"os"
	"regexp"
	"strconv"
)

type CrawlController struct {
	beego.Controller
}

func (c *CrawlController) Crawl() {
	var food models.Food_info
	number := 1
	url := "https://list.tmall.com/search_product.htm?q=%D0%A1%C1%E3%CA%B3&type=p&spm=a220m.1000858.a2227oh.d100&from=.list.pc_1_searchbutton"
	html, err := models.HttpGet(url)
	result := models.ConvertToString(html,"gbk","utf-8")
	models.Check(err)
	for i := 0 ; i < 12*5 ; i++ {
		if i < 5 {
			//正则
			ret1 := regexp.MustCompile(`<a class="productShop-name".*?target="_blank">([\s\S]*?)</a>`)
			ret2 := regexp.MustCompile(`<img  src=  "//img.alicdn.com/bao/uploaded/(.*?)" />`)
			ret3 := regexp.MustCompile(`<a href=".*?" target="_blank" title="(.*?)" data-p=".*?" >`)
			ret4 := regexp.MustCompile(`<em title=".*?"><b>&yen;</b>(.*?)</em>`)
			//提取
			store := ret1.FindAllStringSubmatch(result, -1)
			Pic   := ret2.FindAllStringSubmatch(result, -1)
			intro := ret3.FindAllStringSubmatch(result, -1)
			price := ret4.FindAllStringSubmatch(result, -1)
			//赋值
			food.Id		   = i + 1
			food.FoodStore = store[i][1]
			food.FoodIntro = intro[i][1]
			food.FoodPrice = price[i][1]
			food.FoodPic = "https://img.alicdn.com/bao/uploaded/" + Pic[i][1]
			//下载图片
			imgPath := "C:\\Users\\WIN10\\go\\src\\food_web\\static\\img\\"
			imgUrl := "https://img.alicdn.com/bao/uploaded/" + Pic[i][1]
			filename := imgPath + strconv.Itoa(number) + ".jpg"
			req := httplib.Get(imgUrl)
			content, _ := req.Bytes()
			file, _ := os.Create(filename)
			file.Write(content)
			number ++
			models.AddFood(&food)
		}else{
			//正则
			ret1 := regexp.MustCompile(`<a class="productShop-name".*?target="_blank">([\s\S]*?)</a>`)
			ret2 := regexp.MustCompile(`<img  data-ks-lazyload=  "//img.alicdn.com/bao/uploaded/(.*?)" />`)
			ret3 := regexp.MustCompile(`<a href=".*?" target="_blank" title="(.*?)" data-p=".*?" >`)
			ret4 := regexp.MustCompile(`<em title=".*?"><b>&yen;</b>(.*?)</em>`)
			//提取
			store := ret1.FindAllStringSubmatch(result, -1)
			intro := ret3.FindAllStringSubmatch(result, -1)
			price := ret4.FindAllStringSubmatch(result, -1)
			Pic   := ret2.FindAllStringSubmatch(result, -1)
			//赋值
			food.Id		   = i + 1
			food.FoodStore = store[i][1]
			food.FoodIntro = intro[i][1]
			food.FoodPrice = price[i][1]
			food.FoodPic = "https://img.alicdn.com/bao/uploaded/" + Pic[i-5][1]
			//下载图片
			imgPath := "C:\\Users\\WIN10\\go\\src\\food_web\\static\\img\\"
			imgUrl := "https://img.alicdn.com/bao/uploaded/" + Pic[i-5][1]
			filename := imgPath + strconv.Itoa(number) + ".jpg"
			req := httplib.Get(imgUrl)
			content, _ := req.Bytes()
			file, _ := os.Create(filename)
			file.Write(content)
			number ++
			models.AddFood(&food)
		}
		fmt.Printf("爬取第%d个商品信息已完成\n",i+1)
	}
	c.Ctx.WriteString("爬取商品信息已完成...",)
}
