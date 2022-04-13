package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CPU阈值设置请求体
type SetCpuThresholdData struct {
	// CPU阈值设定值，单位为百分比(%)。 取值范围：0 - 100。

	Cpu *int32 `json:"cpu,omitempty"`
}

func (o SetCpuThresholdData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetCpuThresholdData struct{}"
	}

	return strings.Join([]string{"SetCpuThresholdData", string(data)}, " ")
}
