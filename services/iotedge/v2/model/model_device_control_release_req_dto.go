package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceControlReleaseReqDto 设备控制设置请求结构体
type DeviceControlReleaseReqDto struct {

	// 控制id
	ControlId string `json:"control_id"`

	// 服务id，可选
	ServiceId *string `json:"service_id,omitempty"`

	// 调度计划优先级。
	Priority int32 `json:"priority"`

	// 控制释放的属性数组
	Properties []string `json:"properties"`
}

func (o DeviceControlReleaseReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceControlReleaseReqDto struct{}"
	}

	return strings.Join([]string{"DeviceControlReleaseReqDto", string(data)}, " ")
}
