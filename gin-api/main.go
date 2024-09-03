package main

import (
	"gin-api/global"
	"gin-api/utils"
)

func main() {
	global.Log = utils.InitLogger()
	global.Log.Info("hello world")
	global.Log.Error("parse time error")

}
