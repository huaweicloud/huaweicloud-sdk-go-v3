package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventCreateReq 系统订阅创建请求体
type EventCreateReq struct {
	Systemevent *EventCreateDetail `json:"systemevent,omitempty"`
}

func (o EventCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventCreateReq struct{}"
	}

	return strings.Join([]string{"EventCreateReq", string(data)}, " ")
}
