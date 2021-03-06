package utils

import (
	"github.com/astaxie/beego/context"
	"net"
	"strings"
)

// 服务端获取真实IP地址
func GetIP(ctx *context.Context) string {
	var addr = ""
	addr = ctx.Input.Header("X-Real-IP")
	if addr != "" {
		return addr
	}

	ips := getIps(ctx)
	if len(ips)> 0 && ips[0] !=""{
		ip,_,err:=net.SplitHostPort(ips[0])
		if err!=nil{
			ip = ips[0]
		}
		return ip
	}
	addr = ctx.Request.RemoteAddr
	if ip,_,err:=net.SplitHostPort(addr);err==nil{
		return ip
	}
	return ctx.Request.RemoteAddr
}

// IPs 获取 IP 数组，每经过一级代理(匿名代理除外)，代理服务器都会把这次请求的来源IP数组中
func getIps(ctx *context.Context) []string {
	if ips := ctx.Input.Header("X-Forwarded-For"); ips != "" {
		return strings.Split(ips,",")
	} else {
		return []string{}
	}
}

// 获取请求头中的经纬度
func GetCoordinate(ctx *context.Context) []string {
	coordinate := make([]string, 0)
	lng := ctx.Input.Header("X-Lng")
	lat := ctx.Input.Header("X-Lat")
	if lng != "" && lat != "" {
		coordinate = append(coordinate, lng)
		coordinate = append(coordinate, lat)
	}
	return coordinate
}
