苹果支付

## 鸣谢

[![jetbrains.svg](jetbrains.svg)](https://www.jetbrains.com/?from=AliPay%20SDK%20for%20Go)

## 安装

```go
go get github.com/smartwalle/apple
```

```go
import github.com/smartwalle/apple
```

## 帮助

在集成的过程中有遇到问题，欢迎加 QQ 群 203357977 讨论。

## 其它支付

支付宝 [https://github.com/smartwalle/alipay](https://github.com/smartwalle/alipay)

PayPal [https://github.com/smartwalle/paypal](https://github.com/smartwalle/paypal)

银联支付 [https://github.com/smartwalle/unionpay](https://github.com/smartwalle/unionpay)

## 苹果内购验证

```go
var summary, info, err = apple.VerifyReceipt(transactionId, receipt)
```

苹果内购验证支持**生产环境**和**沙箱环境**，**VerifyReceipt()** 函数内部会优先向苹果生产环境进行验证，然后根据获取到的数据判断是否要向沙箱环境进行验证。

可以从 **VerifyReceipt()** 函数返回的数据中判断该支付所属的环境信息。

## 苹果登录数据解析

```go
var client = apple.NewAuthClient()
var user, err = client.DecodeToken("从客户端获取到的 IdentityToken")
```

## 苹果登录数据验证

如果要验证 Token 的合法性，在初始化 IdentityClient 的时候，需要设置 BundleId。

```go
var client = apple.NewAuthClient(apple.WithBundleId("bundle id"))
var user, err = client.VerifyToken("从客户端获取到的 IdentityToken")
```

## 通知数据解析

```go
var notification, err = apple.DecodeNotification([]byte(data))
```

业务服务器提供一个请求方法为 **POST** 的 HTTP 接口给苹果，苹果会在需要的时候推送一些通知消息到该接口。

```go
var s = gin.Default()
s.POST("/apple", apple)

func apple(c *gin.Context) {
    var data, _ = io.ReadAll(c.Request.Body)
    var notification, err = apple.DecodeNotification([]byte(data)) 
    // 关于这里如何返回数据参考 https://developer.apple.com/documentation/appstoreservernotifications/responding_to_app_store_server_notifications
    // 简单来讲，返回 HTTP Status Code 200 表示我们成功处理该通知
    // 如：c.Status(http.StatusOK)
	
    // 返回 HTTP Status Code 50x 或者 40x 表示我们没有成功处理该通知，苹果会在一定时间后重新推送该通知
    // 如：c.Status(http.StatusBadRequest)
}

```

## 其它接口

* **[Look Up Order ID](https://developer.apple.com/documentation/appstoreserverapi/look_up_order_id)**
* **[Get Refund History](https://developer.apple.com/documentation/appstoreserverapi/get_refund_history)**
* **[Get All Subscription Statuses](https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses)**
* **[Extend a Subscription Renewal Date](https://developer.apple.com/documentation/appstoreserverapi/extend_a_subscription_renewal_date)**
* **[Get Transaction History](https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history)**
* **[Send Consumption Information](https://developer.apple.com/documentation/appstoreserverapi/send_consumption_information)**

以上接口需要先初始化 apple.Client

```go
var client, _ = apple.New(keyfile, keyId, issuer, bundleId, isProduction)
```
#### 关于 keyfile, keyId, issuer 如何获取？

[Creating API Keys to Use With the App Store Server API
](https://developer.apple.com/documentation/appstoreserverapi/creating_api_keys_to_use_with_the_app_store_server_api)

## License

This project is licensed under the MIT License.
