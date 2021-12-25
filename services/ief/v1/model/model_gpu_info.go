package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GpuInfo struct {
	// GPU名称

	Name *string `json:"name,omitempty"`
	// GPU类型

	Type *string `json:"type,omitempty"`
	// GPU memory大小，单位MB

	Capacity *string `json:"capacity,omitempty"`
}

func (o GpuInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GpuInfo struct{}"
	}

	return strings.Join([]string{"GpuInfo", string(data)}, " ")
}
