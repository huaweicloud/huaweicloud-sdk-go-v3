package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EngineSpec 微服务引擎相关参数。
type EngineSpec struct {

	// CPU及内存规格。
	AvailableCpuMemory *string `json:"availableCpuMemory,omitempty"`

	// 是否为线性规格。
	Linear *string `json:"linear,omitempty"`

	// 可用节点规格类型前缀。
	AvailablePrefix *string `json:"availablePrefix,omitempty"`
}

func (o EngineSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineSpec struct{}"
	}

	return strings.Join([]string{"EngineSpec", string(data)}, " ")
}
