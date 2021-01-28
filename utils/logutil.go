package utils

import (
	"encoding/json"
	"fmt"
	"github.com/hyperknob/logger"
	"github.com/astaxie/beego"
)

/**
 * 为logStash记录的Beego接口日志
 */
func RecordLog4LogStash(methodType string, body interface{}, c beego.Controller) {
	recordVal := make(map[string]interface{})
	recordVal["controllerName"], recordVal["actionName"] = c.GetControllerAndAction()
	recordVal["actionType"] = methodType
	recordVal["protocol"] = c.Ctx.Request.Proto
	recordVal["routePath"] = c.Ctx.Request.RequestURI
	if methodType == "req" {
		recordVal["header"] = c.Ctx.Request.Header
		recordVal["remoteAddr"] = c.Ctx.Request.RemoteAddr
	} else if methodType == "res" {
		c.Ctx.Output.Header("X-Request-Id", c.Ctx.Input.Header("X-Request-Id"))
		recordVal["header"] = c.Ctx.ResponseWriter.Header()
	}
	recordVal["body"] = body
	recordValBytes, _ := json.Marshal(recordVal)
	logger.Info(string(recordValBytes))
}

/**
 * 为logStash记录的普通日志
 */
func Info(pattern string, variables... interface{}) {
	recordVal := make(map[string]interface{})
	recordVal["body"] = fmt.Sprintf(pattern, variables)
	recordValBytes, _ := json.Marshal(recordVal)
	logger.Info(string(recordValBytes))
}

/**
 * 为logStash记录的警告日志
 */
func Warn(pattern string, variables... interface{}) {
	recordVal := make(map[string]interface{})
	recordVal["body"] = fmt.Sprintf(pattern, variables)
	recordValBytes, _ := json.Marshal(recordVal)
	logger.Warn(string(recordValBytes))
}

/**
 * 为logStash记录的错误日志
 */
func Error(pattern string, variables... interface{}) {
	recordVal := make(map[string]interface{})
	recordVal["body"] = fmt.Sprintf(pattern, variables)
	recordValBytes, _ := json.Marshal(recordVal)
	logger.Error(string(recordValBytes))
}