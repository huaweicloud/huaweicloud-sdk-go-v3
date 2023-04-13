package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComputeSpec struct {

	// 算力规格类型，如C6
	FlavorType string `json:"flavor_type"`

	// 计算单元设备数
	Count int32 `json:"count"`
}

func (o ComputeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeSpec struct{}"
	}

	return strings.Join([]string{"ComputeSpec", string(data)}, " ")
}
