package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomizeSourceCreateReq struct {

	// 自定义事件源名称，租户下唯一，由小写字母、数字、点、下划线和中划线组成，必须以字母或数字开头，且不能以hc.开头
	Name string `json:"name"`

	// 事件源描述
	Description *string `json:"description,omitempty"`

	// 指定事件源归属的事件通道ID
	ChannelId string `json:"channel_id"`
}

func (o CustomizeSourceCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeSourceCreateReq struct{}"
	}

	return strings.Join([]string{"CustomizeSourceCreateReq", string(data)}, " ")
}
