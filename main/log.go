package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "info":
		return logs.LevelInfo
	case "warn":
		return logs.LevelWarn
	case "error":
		return logs.LevelError
	case "critical":
		return logs.LevelCritical
	}
	return logs.LevelDebug
}

func initLogger(path, level string) error {
	config := make(map[string]interface{})
	config["filename"] = path
	config["level"] = convertLogLevel(level)
	bs, err := json.Marshal(&config)
	if err != nil {
		fmt.Println("initLogger failed, marshl1 err:", err)
		return err
	}
	logs.SetLogger(logs.AdapterFile, string(bs))
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	return nil
}
