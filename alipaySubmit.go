/* *
 *类名：AlipaySubmit
 *功能：支付宝各接口请求提交类
 *详细：构造支付宝各接口表单HTML文本，获取远程HTTP数据
 *版本：3.3
 *日期：2012-08-13
 *说明：
 *以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写,并非一定要使用该代码。
 *该代码仅供学习和研究支付宝接口使用，只是提供一个参考。
 */
package alipay

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

/**
 * 支付宝提供给商户的服务接入网关URL(新)
 */
var ALIPAY_GATEWAY_NEW = "https://mapi.alipay.com/gateway.do?"

var alipayConfig = AlipayConfig{
	InputCharset: "utf-8",
	SignType:     "MD5",
}

/*初始化支付宝参数*/
func InitAlipayConfig(partner, sellerMmail, key, returnUrl, notifyUrl string) {
	alipayConfig.Partner = partner
	alipayConfig.SellerMmail = sellerMmail
	alipayConfig.Key = key
	alipayConfig.ReturnUrl = returnUrl
	alipayConfig.NotifyUrl = notifyUrl
}

/**
 * 生成签名结果
 * @param sPara 要签名的数组
 * @return 签名结果字符串
 */
func BuildRequestMysign(sPara map[string]string) string {

	sParaStr := CreateLinkString(sPara)
	return Sign(sParaStr, alipayConfig.Key)
}

/**
 * 生成要请求给支付宝的参数数组
 * @param sParaTemp 请求前的参数数组
 * @return 要请求的参数数组
 */
func BuildRequestPara(sParaTemp map[string]string) map[string]string {
	ret := ParaFilter(sParaTemp)
	ret["sign"] = BuildRequestMysign(ret)
	ret["sign_type"] = alipayConfig.SignType
	return ret
}

/**
 * 建立请求，以表单HTML形式构造（默认）
 * @param sParaTemp 请求参数数组
 * @param strMethod 提交方式。两个值可选：post、get
 * @param strButtonName 确认按钮显示文字(不要了，基本没用)
 * @return 提交表单HTML文本
 */
func BuildRequest(sParaTemp map[string]string, strMethod string) string {

	sPara := BuildRequestPara(sParaTemp)
	html := bytes.Buffer{}
	html.WriteString(fmt.Sprintf("<form id=\"alipaysubmit\" name=\"alipaysubmit\" action=\"%s_input_charset=%s\" method=\"%s\" style=\"display:none;\">", ALIPAY_GATEWAY_NEW, alipayConfig.InputCharset, strMethod))
	ms := NewMapSorter(sPara)
	for _, item := range ms {
		html.WriteString(fmt.Sprintf("<input type=\"hidden\" name=\"%s\" value=\"%s\"/>", item.Key, item.Value))
	}
	html.WriteString("<input type=\"submit\" value=\"提交\" style=\"display:none;\"></form><script>document.forms['alipaysubmit'].submit();</script>")
	return html.String()

}

/**
默认的支付功能，主要为提供即时到账使用，帮添加默认的请求参数，可以自行掉BuildRequest方法构造支付表单
@orderId 商品在商家平台上的ID
@fee 交易金额
@subject 交易说明
@paymethod 支付方式，bankPay为网银支付，directPay 余额支付
@defaultbank 银行简码，若为余额支付时可为空
*/
func AlipayToPay(orderId string, fee float64, subject, paymethod, defaultbank string) (string, error) {

	if paymethod != "bankPay" && paymethod != "directPay" {
		return "", errors.New("paymethod params is error")
	}
	params := map[string]string{}
	params["body"] = subject + ",交易金额：" + strconv.FormatFloat(fee, 'f', 2, 32) + "元"
	params["_input_charset"] = alipayConfig.InputCharset
	params["notify_url"] = alipayConfig.NotifyUrl
	params["out_trade_no"] = orderId
	params["partner"] = alipayConfig.Partner
	params["payment_type"] = "1"
	params["return_url"] = alipayConfig.ReturnUrl
	params["seller_email"] = alipayConfig.SellerMmail
	params["service"] = "create_direct_pay_by_user"
	params["subject"] = subject
	params["total_fee"] = strconv.FormatFloat(float64(fee), 'f', 2, 32)
	if paymethod == "bankPay" && len(defaultbank) != 0 { //网银支付
		params["defaultbank"] = defaultbank
	}
	params["paymethod"] = paymethod
	return BuildRequest(params, "get"), nil
}
