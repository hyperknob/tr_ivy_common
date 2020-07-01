package tr_ivy_common

import "github.com/gbrlsnchs/jwt/v3"

// 定制化的jwt payload
type CustomJWTPayload struct {
	jwt.Payload
	OpenID      string `json:"openid,omitempty"`      // 第三方用户唯一标识
	OpenChannel string `json:"openchannel,omitempty"` // 第三方用户渠道
	UID         string `json:"uid,omitempty"`         // 用户唯一标识
}




