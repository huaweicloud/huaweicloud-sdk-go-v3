package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomizeSourceUpdateReq struct {

	// 事件源描述
	Description *string `json:"description,omitempty"`

	// json格式封装消息实例更新信息：如RabbitMQ实例的虚拟主机vhost字段、队列queue字段、用户密码
	Detail *interface{} `json:"detail,omitempty"`
}

func (o CustomizeSourceUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeSourceUpdateReq struct{}"
	}

	return strings.Join([]string{"CustomizeSourceUpdateReq", string(data)}, " ")
}
