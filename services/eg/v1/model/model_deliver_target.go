package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeliverTarget struct {

	// 投递目标ID，即事件目标ID
	DeliverTargetId *string `json:"deliverTargetId,omitempty"`

	// 投递目标名称，及事件目标名称
	DeliverTargetName *string `json:"deliverTargetName,omitempty"`

	// 投递状态         SUCCESS Or  FAILED
	DeliverStatus *string `json:"deliverStatus,omitempty"`

	// 考虑展示的个数    例如限制只展示最新三条
	DeliverDetailList *[]DeliverDetail `json:"deliverDetailList,omitempty"`
}

func (o DeliverTarget) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeliverTarget struct{}"
	}

	return strings.Join([]string{"DeliverTarget", string(data)}, " ")
}
