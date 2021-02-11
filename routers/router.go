package routers

import (
    "food_web/controllers"
    "github.com/astaxie/beego"
)

func init() {
    //爬虫路由
    beego.Router("/crawl",&controllers.CrawlController{},"*:Crawl")

    //主页路由
    beego.Router("/",&controllers.HomeController{})
    //注册路由
    beego.Router("/register",&controllers.RegisterController{})
    //登录路由
    beego.Router("/login",&controllers.LoginController{})
    //退出路由
    beego.Router("/exit",&controllers.ExitController{})
    //个人页面
    beego.Router("/selfhome",&controllers.SelfHomeController{})
    //修改信息
    beego.Router("/alter",&controllers.AlterController{})
    //头像文件上传
    beego.Router("/head",&controllers.HeadController{})
    //修改密码
    beego.Router("/change",&controllers.AltPasswordController{})
    //加入购物车
    beego.Router("/collect",&controllers.CollectController{})
    //移出购物车
    beego.Router("/delete/:id",&controllers.DeleteCarController{})
    //商品详细页
    beego.Router("/food/:id",&controllers.DetailController{})
    //购物车显示
    beego.Router("/shop_car",&controllers.ShopcarController{})
    //商家上架商品
    beego.Router("/food/add",&controllers.FoodLoadController{})
    //商家文件上传
    beego.Router("/upload",&controllers.UploadController{})
}
