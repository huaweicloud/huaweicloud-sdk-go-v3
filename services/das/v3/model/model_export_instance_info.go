package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportInstanceInfo 导出实例信息
type ExportInstanceInfo struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例状态
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 指标信息
	Metrics map[string]float64 `json:"metrics,omitempty"`
}

func (o ExportInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportInstanceInfo struct{}"
	}

	return strings.Join([]string{"ExportInstanceInfo", string(data)}, " ")
}
