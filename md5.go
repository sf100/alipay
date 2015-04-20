package alipay

import (
	"crypto/md5"
	"encoding/hex"
)

/**
 * 生成签名结果
 * @param sPara 要签名的数组
 * @return 签名结果字符串
 */
func Sign(text, key string) string {
	text = text + key
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
