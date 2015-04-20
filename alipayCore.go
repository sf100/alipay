/** 该sdk都是从 支付宝官方java版SDK翻译而来
 *功能：支付宝接口公用函数类
 *详细：该类是请求、通知返回两个文件所调用的公用函数核心处理文件，不需要修改
 *说明：
 *以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写,并非一定要使用该代码。
 *该代码仅供学习和研究支付宝接口使用，只是提供一个参考。
 */
package alipay

import (
	"bytes"
	"sort"
)

/* *
 *类名：AlipayConfig
 *功能：基础配置类
 */
type AlipayConfig struct {
	Partner      string
	SellerMmail  string
	Key          string
	InputCharset string
	SignType     string
	NotifyUrl    string
	ReturnUrl    string
}

/*为了构造有序Map*/
type MapSorter []Item

func (ms MapSorter) Len() int {
	return len(ms)
}

func (ms MapSorter) Less(i, j int) bool {
	return ms[i].Key < ms[j].Key // 按键排序
}

func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

type Item struct {
	Key   string
	Value string
}

/*构造有序Map*/
func NewMapSorter(m map[string]string) MapSorter {

	ms := make(MapSorter, 0, len(m))

	for k, v := range m {
		ms = append(ms, Item{k, v})
	}
	sort.Sort(ms)
	return ms
}

/**
 * 除去数组中的空值和签名参数
 * @param sArray 签名参数组
 * @return 去掉空值与签名参数后的新签名参数组
 */
func ParaFilter(sArray map[string]string) map[string]string {

	var resultSArray = make(map[string]string, len(sArray))

	for k, v := range sArray {

		if len(v) == 0 || k == "sign" || k == "sign_type" {
			continue
		}
		resultSArray[k] = v
	}
	return resultSArray
}

/**
 * 把数组所有元素排序，并按照“参数=参数值”的模式用“&”字符拼接成字符串
 * @param params 需要a-z排序并参与字符拼接的参数组
 * @return 拼接后字符串
 */
func CreateLinkString(params map[string]string) string {

	ret := bytes.Buffer{}
	ms := NewMapSorter(params)
	l := len(ms) - 1
	for i, item := range ms {
		if i < l {
			ret.WriteString(item.Key + "=" + item.Value + "&")
		} else {
			ret.WriteString(item.Key + "=" + item.Value)
		}
	}
	return ret.String()
}
