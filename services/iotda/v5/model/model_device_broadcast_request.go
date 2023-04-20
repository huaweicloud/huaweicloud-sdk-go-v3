package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeviceBroadcastRequest struct {

	// **参数说明**：资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，建议携带该参数指定广播消息所属的资源空间，否则广播消息将会向[[默认资源空间](https://support.huaweicloud.com/usermanual-iothub/iot_01_0006.html#section0)](tag:hws)[[默认资源空间](https://support.huaweicloud.com/intl/zh-cn/usermanual-iothub/iot_01_0006.html#section0)](tag:hws_hk)下订阅指定topic的设备发送。 **取值范围**：长度不超过36，只允许字母、数字、下划线（_）、连接符（-）的组合。
	AppId *string `json:"app_id,omitempty"`

	// **参数说明**：接收广播消息的完整Topic名称, 必选。用户需要发布广播消息给设备时，可以使用该参数指定完整的topic名称，物联网平台会向指定资源空间下订阅了该topic的所有在线设备发送消息。广播的topic无需控制台创建，但topic的前缀必须为$oc/broadcast/
	TopicFullName string `json:"topic_full_name"`

	// **参数说明**：广播消息的内容，您需要将消息原文转换成二进制数据并进行Base64编码，Base64编码后的长度不超过128KB。
	Message string `json:"message"`
}

func (o DeviceBroadcastRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceBroadcastRequest struct{}"
	}

	return strings.Join([]string{"DeviceBroadcastRequest", string(data)}, " ")
}
