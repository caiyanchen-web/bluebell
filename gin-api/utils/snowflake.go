package utils

import (
	"gin-api/config"
	"gin-api/global"
	sf "github.com/bwmarrin/snowflake"
	"time"
)

// 雪花算法生成用户id
// 初始化全局node节点
var node *sf.Node

func init() {
	var (
		startTime string    // 开始时间
		st        time.Time // 开始时间
		machineID int64     // 机器id
		err       error     // 错误
	)
	// 设置开始时间，默认为项目创建时间
	startTime = time.Now().Format("2006-01-02")
	//st获取开始时间，年-月-日格式
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		global.Log.Error("parse time error", err)
	}
	machineID = config.Config.System.MachineID
	// 获取时间因子
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
