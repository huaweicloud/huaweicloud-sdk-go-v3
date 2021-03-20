package model

import (
	"encoding/json"

	"strings"
)

// 创建订阅请求结构体
type CreateSubReq struct {
	Subject *Subject `json:"subject"`
	// 订阅的回调地址，用于接收对应资源事件的通知消息，例如：https://10.10.10.10:443/callbackurltest。

	Callbackurl string `json:"callbackurl"`
	// 资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，建议携带该参数指定订阅哪个资源空间下的消息通知，否则订阅的消息通知将会归属到[默认资源空间](https://support.huaweicloud.com/usermanual-iothub/iot_01_0006.html#section0)下。

	AppId *string `json:"app_id,omitempty"`
	// 物联网平台推送通知消息时使用的协议通道。使用“http”填充，表示该订阅推送协议通道为http(s)协议。

	Channel string `json:"channel"`
}

func (o CreateSubReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSubReq struct{}"
	}

	return strings.Join([]string{"CreateSubReq", string(data)}, " ")
}
