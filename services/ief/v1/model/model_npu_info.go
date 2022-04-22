package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NpuInfo struct {

	// NPU名称
	Name *string `json:"name,omitempty"`

	// NPU类型
	Type *string `json:"type,omitempty"`

	// NPU memory大小
	Capacity *string `json:"capacity,omitempty"`

	// NPU驱动版本
	DriverVersion *string `json:"driver_version,omitempty"`
}

func (o NpuInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NpuInfo struct{}"
	}

	return strings.Join([]string{"NpuInfo", string(data)}, " ")
}
