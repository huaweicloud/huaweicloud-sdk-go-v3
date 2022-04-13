package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SendCommandRequestBody struct {
	// 服务ID，自动向下取整

	ServiceId int32 `json:"service_id"`
	// 命令ID，自动向下取整

	CommandId int32 `json:"command_id"`
	// 命令是否同步 true-同步 false-异步 同步命令会将命令以MQTT消息发送给设备后等待设备的MQTT命令响应，收到响应后再回复响应（默认超时5秒），异步命令则以立即返回响应

	IsSync bool `json:"is_sync"`
	// 请求参数列表

	Requests *[]RequestParameter `json:"requests,omitempty"`
}

func (o SendCommandRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendCommandRequestBody struct{}"
	}

	return strings.Join([]string{"SendCommandRequestBody", string(data)}, " ")
}
