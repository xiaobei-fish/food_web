# food_web

欠缺点：

可以在注册controller里面加上插入数据库该用户的金额，初始值为0

然后在购买的时候根据food_id去查询原商家，然后更新商家的金额，但是初始的时候本来都上架了货品还是空头的，会出现一点差错就没有写这个功能

早期没有考虑并发性，用Redis代替了数据库进行一部分数据的储存取用，要考虑并发性将其改成数据库就行。主要还是自己前端JS不会写，不然不需要这些东西、



### conf/app.conf 配置

```shell
appname = food_web
httpport = 8090
runmode = dev

#mysql配置
driverName = mysql
mysqluser = 账号
mysqlpwd = 密码
host = localhost
port = 3306
dbname = 数据库

#session
Sessionon = true
sessionprovider = "file"
sessionname = "web"
sessiongcmaxlifetime = 5400
sessionproviderconfig = "./tmp"
sessioncookielifetime = 5400

#页码配置
foodListPageNum = 5
userListPageNum = 20
```

