package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceControlSetReqDto 设备控制设置请求结构体
type DeviceControlSetReqDto struct {

	// 控制id
	ControlId string `json:"control_id"`

	// 服务id，可选
	ServiceId *string `json:"service_id,omitempty"`

	// 调度计划优先级。
	Priority int32 `json:"priority"`

	// 控制结束时间，毫秒级时间戳
	EndTime *int64 `json:"end_time,omitempty"`

	// 属性key和value的map，用于设置属性的值
	Properties *interface{} `json:"properties"`
}

func (o DeviceControlSetReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceControlSetReqDto struct{}"
	}

	return strings.Join([]string{"DeviceControlSetReqDto", string(data)}, " ")
}
