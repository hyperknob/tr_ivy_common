/**
 * @Author: jinsong
 * @Description:
 * @File:  base64util
 * @Version: 1.0.0
 * @Date: 2019/11/20 19:02
 */
package utils

import (
	"encoding/base64"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var coder = base64.NewEncoding(base64Table)

func Base64Encode(encode_byte []byte) []byte {
	return []byte(coder.EncodeToString(encode_byte))
}

func Base64Decode(decode_byte []byte) ([]byte, error) {
	return coder.DecodeString(string(decode_byte))
}