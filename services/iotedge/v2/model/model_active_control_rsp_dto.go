package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActiveControlRspDto 执行中控制响应结构体
type ActiveControlRspDto struct {

	// 控制类型，包括control、schedule、local_control
	ControlType *string `json:"control_type,omitempty"`

	// 控制id
	ControlId *string `json:"control_id,omitempty"`

	// 控制的优先级
	Priority *int32 `json:"priority,omitempty"`

	// 此次控制的值
	Value *interface{} `json:"value,omitempty"`

	// 控制的结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 记录创建的时间
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o ActiveControlRspDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActiveControlRspDto struct{}"
	}

	return strings.Join([]string{"ActiveControlRspDto", string(data)}, " ")
}
