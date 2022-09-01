package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NpuInfo struct {

	// NPU名称
	Name *string `json:"name,omitempty" xml:"name"`

	// NPU类型
	Type *string `json:"type,omitempty" xml:"type"`

	// NPU memory大小
	Capacity *string `json:"capacity,omitempty" xml:"capacity"`

	// NPU驱动版本
	DriverVersion *string `json:"driver_version,omitempty" xml:"driver_version"`
}

func (o NpuInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NpuInfo struct{}"
	}

	return strings.Join([]string{"NpuInfo", string(data)}, " ")
}
