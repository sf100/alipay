# Alipay Go SDK

**默认带有支持即时到账、网银支付的方法**

该 SDK 是从支付宝官方 Java 版 SDK 翻译而来，提供大家参考和学习，因某些方法不常用，目前 alipaySubmit 未编写方法有：

/* 建立请求，以表单HTML形式构造，带文件上传功能 */
`buildRequest(String strParaFileName, String strFilePath,Map<String, String> sParaTemp)`

/* 建立请求，以模拟远程HTTP的POST请求方式构造并获取支付宝的处理结果 */
`buildRequest(Map<String, String> sParaTemp, String strMethod, String strButtonName, String strParaFileName)`

## Examples
  请参考test/testMain.go
  如果AlipayToPay支付方法不满足需求，可根据自己的参数需求调用BuildRequest方法来构造提交支付宝的表单


# 开源协议

该源码使用 [Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0.txt) 开源。
