package main

import (
	"ginblog/model"
	_ "ginblog/routes" //我有个ginblog的文件夹,要先进去，再在名为ginblog的workspace工作
)

func main() {

	//routes.InitRouter()
	model.InitDb()

}
