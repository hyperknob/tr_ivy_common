package utils

import (
	"github.com/shomali11/util/xhashes"
	"strconv"
)
import "github.com/mitchellh/hashstructure"


/**
 * 获取任意结构的hash int值
 */
func GetHashValue(v interface{}) uint64 {
	hash, err := hashstructure.Hash(v, nil)
	if err != nil {
		panic(err)
	}
	return hash
}

/**
 * 获取任意String的FNV32 hash值
 */
func GetStringHashValue(v string) string {
	return strconv.FormatInt(int64(xhashes.FNV32a(v)), 10)
}