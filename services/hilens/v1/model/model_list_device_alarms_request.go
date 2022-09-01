package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDeviceAlarmsRequest struct {

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量，取值范围1~100，默认为100
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 设备ID
	DeviceId *string `json:"device_id,omitempty" xml:"device_id"`
}

func (o ListDeviceAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeviceAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListDeviceAlarmsRequest", string(data)}, " ")
}
