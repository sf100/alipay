/* *
 *类名：AlipayNotify
 *功能：支付宝通知处理类
 *详细：处理支付宝各接口通知返回
 *以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写,并非一定要使用该代码。
 *该代码仅供学习和研究支付宝接口使用，只是提供一个参考

 *************************注意*************************
 *调试通知返回时，可查看或改写log日志的写入TXT里的数据，来检查通知返回是否正常
 */
package alipay

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var HTTPS_VERIFY_URL = "https://mapi.alipay.com/gateway.do?service=notify_verify&"

/**
 * 验证消息是否是支付宝发出的合法消息
 * @param params 通知返回来的参数数组
 * @return 验证结果
 */
func Verify(params map[string]string) bool {

	var responseTxt string
	if len(params["notify_id"]) > 0 {
		notify_id := params["notify_id"]
		responseTxt = VerifyResponse(notify_id)
	} else {
		return false
	}

	sign := params["sign"]
	isSign := GetSignVeryfy(params, sign)
	if isSign && responseTxt == "true" {
		return true
	} else {
		return false
	}

}

/**
 * 根据反馈回来的信息，生成签名结果
 * @param Params 通知返回来的参数数组
 * @param sign 比对的签名结果
 * @return 生成的签名结果
 */
func GetSignVeryfy(params map[string]string, sign string) bool {

	parasNew := ParaFilter(params)
	preSignStr := CreateLinkString(parasNew)
	fmt.Println(params)
	if alipayConfig.SignType == "MD5" {
		newSign := Sign(preSignStr, alipayConfig.Key)
		fmt.Println(newSign)
		fmt.Println(sign)
		if newSign == sign {
			return true
		}
	}
	return false
}

/**
 * 获取远程服务器ATN结果,验证返回URL
 * @param notify_id 通知校验ID
 * @return 服务器ATN结果
 * 验证结果集：
 * invalid命令参数不对 出现这个错误，请检测返回处理中partner和key是否为空
 * true 返回正确信息
 * false 请检查防火墙或者是服务器阻止端口问题以及验证时间是否超过一分钟
 */
func VerifyResponse(notify_id string) string {

	veryfy_url := HTTPS_VERIFY_URL + "partner=" + alipayConfig.Partner + "&notify_id=" + notify_id
	resp, err := http.Post(veryfy_url, "text/plain;charset=UTF-8", nil)
	if err != nil {
		fmt.Println(err)
		return "false"
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "false"
	}
	defer resp.Body.Close()

	return string(body)
}

/** 支付宝同步调用返回接口
status 返回支付结果
orderId 商品订单ID
buyerEmail 买家邮箱
tradeNo 支付宝的订单号
return 不管支付是否成功，都返回支付返回的订单信息，给使用着，方便调用者做记录
*/
func AlipayReturn(request *http.Request) (status bool, orderId, buyerEmail, tradeNo string) {

	params := make(map[string]string, 19)
	request.ParseForm()
	urlValus := request.Form

	for k, v := range urlValus {
		params[k] = strings.Join(v, "")
	}

	if len(params["out_trade_no"]) == 0 {

		return false, "", "", ""

	} else {
		//验证回调请求是否合法
		if Verify(params) {
			//交易成功
			if params["trade_status"] == "TRADE_FINISHED" || params["trade_status"] == "TRADE_SUCCESS" {
				return true, params["out_trade_no"], params["buyer_email"], params["trade_no"]
			} else {
				return false, params["out_trade_no"], params["buyer_email"], params["trade_no"]
			}
		} else {
			return false, params["out_trade_no"], params["buyer_email"], params["trade_no"]
		}
	}

	return
}

/** 支付宝同步调用返回接口
status 返回支付结果
orderId 商品订单ID
buyerEmail 买家邮箱
tradeNo 支付宝的订单号
*/
func AlipayNotify(request *http.Request) (status bool, orderId, buyerEmail, tradeNo string) {

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		status = false
		return
	}

	bodyStr, _ := url.QueryUnescape(string(body))
	postArray := strings.Split(bodyStr, "&")
	params := map[string]string{}
	for _, v := range postArray {
		detail := strings.Split(v, "=")
		params[detail[0]] = detail[1]
	}

	if len(params["out_trade_no"]) == 0 {

		return false, params["out_trade_no"], params["buyer_email"], params["trade_no"]

	} else {
		//验证回调请求是否合法
		if Verify(params) {
			//交易成功
			if params["trade_status"] == "TRADE_FINISHED" || params["trade_status"] == "TRADE_SUCCESS" {
				return true, params["out_trade_no"], params["buyer_email"], params["trade_no"]
			} else {
				status = false
				return
			}
		} else {
			tradeNo = fmt.Sprintln(params)
			status = false
			return
		}
	}

}
