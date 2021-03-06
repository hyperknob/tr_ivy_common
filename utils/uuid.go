package utils

import (
	"github.com/gofrs/uuid"
	"github.com/mr-tron/base58"
)

/*
	高并发情况下，生成全局唯一id
	1、注意幂等性且全局唯一
	2、注意安全性，不能被猜疑破解
	3、趋势递增型
	比如订单号 ：业务编码 + 时间戳 + 机器编码（前四位） + 随机数（四位） + 毫秒
    采用uuid1+uuid4并压缩连接构成, 和python管理后台一致
*/

func GenUuid() string {
	/*
		uuid 组成部分：当前日期和时间 + 时钟序列 + 随机数 + 全球唯一的IEEE识别码
		全球唯一的IEEE识别码 ：从网卡mac 中获得，没有网卡尝试从其他地方获得
	*/
	// 创建
	id1, _ := uuid.NewV1() // based on timestamp and MAC address (RFC 4122)
	id4, _ := uuid.NewV4() // Version 4, based on random numbers (RFC 4122)
	return base58.Encode(id1.Bytes()) + base58.Encode(id4.Bytes())
}
