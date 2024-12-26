package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomizeSourceUpdateReq struct {

	// 事件源描述
	Description *string `json:"description,omitempty"`

	Detail *RocketMqDetail `json:"detail,omitempty"`
}

func (o CustomizeSourceUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeSourceUpdateReq struct{}"
	}

	return strings.Join([]string{"CustomizeSourceUpdateReq", string(data)}, " ")
}
