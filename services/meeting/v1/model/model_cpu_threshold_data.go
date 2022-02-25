package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CPU阈值查询结果结构体
type CpuThresholdData struct {
	// 自定义的cpu阈值，单位为百分比(%)

	Cpu *int32 `json:"cpu,omitempty"`
	// cpu阈值默认值，单位为百分比(%)

	CpuDefault *int32 `json:"cpuDefault,omitempty"`
}

func (o CpuThresholdData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CpuThresholdData struct{}"
	}

	return strings.Join([]string{"CpuThresholdData", string(data)}, " ")
}
