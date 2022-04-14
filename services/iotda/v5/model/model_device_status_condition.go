package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 条件中设备状态类型的信息，自定义结构。
type DeviceStatusCondition struct {
	// **参数说明**：状态列表，设备状态条件携带该参数。

	StatusList *[]string `json:"status_list,omitempty"`
}

func (o DeviceStatusCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceStatusCondition struct{}"
	}

	return strings.Join([]string{"DeviceStatusCondition", string(data)}, " ")
}
