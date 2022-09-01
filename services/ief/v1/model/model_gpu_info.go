package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GpuInfo struct {

	// GPU名称
	Name *string `json:"name,omitempty" xml:"name"`

	// GPU类型
	Type *string `json:"type,omitempty" xml:"type"`

	// GPU memory大小，单位MB
	Capacity *string `json:"capacity,omitempty" xml:"capacity"`
}

func (o GpuInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GpuInfo struct{}"
	}

	return strings.Join([]string{"GpuInfo", string(data)}, " ")
}
