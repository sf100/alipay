# Alipay Go SDK

**默认带有支持即时到账、网银支付的方法**

该 SDK 是从支付宝官方 Java 版 SDK 翻译而来，提供大家参考和学习，因某些方法不常用，目前 alipaySubmit 未编写方法有：

/* 建立请求，以表单HTML形式构造，带文件上传功能 */
`buildRequest(String strParaFileName, String strFilePath,Map<String, String> sParaTemp)`

/* 建立请求，以模拟远程HTTP的POST请求方式构造并获取支付宝的处理结果 */
`buildRequest(Map<String, String> sParaTemp, String strMethod, String strButtonName, String strParaFileName)`

example
======

func init() {
	alipay.InitAlipayConfig("test", "test@test.com", "test", "http://127.0.0.1:8089/alipay/return", "http://127.0.0.1:8089/alipay/notify")
}
func main() {

	http.HandleFunc("/pay", pay)
	http.HandleFunc("/alipay/return", alipayReturn)
	http.HandleFunc("/alipay/notify", alipayNotify)

	err := http.ListenAndServe(":8089", nil)
	CheckErr(err)

}

func pay(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("pay.html")
	CheckErr(err)

	/*
	  CMB为银行简码，银行简码可以参考支付宝"纯网关接口(create_direct_pay_by_user).pdf"文档10.6章节,如果为即时到账银行简码可为空
	  "bankPay"为网银支付，余额支付"directPay"
	*/
	form, err := alipay.AlipayToPay("orderID", 0.01, "用户充值测试", "bankPay", "CMB")
	CheckErr(err)
	formHtml := template.HTML(form)
	t.Execute(w, formHtml)

}

/*页面跳转同步通知*/
func alipayReturn(w http.ResponseWriter, r *http.Request) {

	/**
	  status 返回支付结果
	  orderId 商品订单ID
	  buyerEmail 买家邮箱
	  tradeNo 支付宝的订单号
	*/
	status, orderId, buyerEmail, tradeNo := alipay.AlipayReturn(r)
	if status {
		fmt.Printf("商品订单号：%s , 买家账号：%s,支付宝订单号：%s", orderId, buyerEmail, tradeNo)
	} else {
		w.Write([]byte("异常错误！"))
	}
}

/*服务器异步通知*/
func alipayNotify(w http.ResponseWriter, r *http.Request) {

	/**
	  status 返回支付结果
	  orderId 商品订单ID
	  buyerEmail 买家邮箱
	  tradeNo 支付宝的订单号
	*/
	status, orderId, buyerEmail, tradeNo := alipay.AlipayNotify(r)
	if status {
		fmt.Printf("商品订单号：%s , 买家账号：%s,支付宝订单号：%s", orderId, buyerEmail, tradeNo)
	} else {
		w.Write([]byte("异常错误！"))
	}
}


# 开源协议

该源码使用 [Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0.txt) 开源。
