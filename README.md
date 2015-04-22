# alipay
**默认带有支持即时到账、网银支付的方法**

该sdk都是从支付宝官方java版SDK翻译而来,提供大家参考和学习，因某些方法不常用，目前alipaySubmit 未编写方法有：

/*建立请求，以表单HTML形式构造，带文件上传功能*/

````buildRequest(String strParaFileName, String strFilePath,Map<String, String> sParaTemp)```

/* 建立请求，以模拟远程HTTP的POST请求方式构造并获取支付宝的处理结果*/

````buildRequest(Map<String, String> sParaTemp, String strMethod, String strButtonName, String strParaFileName)```
