package utils

import (
	"encoding/json"
	"github.com/wonderivan/logger"
)

/**
 * 为logStash记录的Beego接口日志
 */
func RecordLog4LogStash(methodName string, methodType string, methodProtocol, input interface{}) {
	recordVal := make(map[string]interface{})
	recordVal["methodName"] = methodName
	recordVal["methodType"] = methodType
	recordVal["methodProtocol"] = methodProtocol
	recordVal["parameters"] = input
	recordValBytes, _ := json.Marshal(recordVal)
	logger.Info(string(recordValBytes))
}