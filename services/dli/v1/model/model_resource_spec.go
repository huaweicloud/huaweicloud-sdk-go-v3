package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceSpec 资源规格。
type ResourceSpec struct {

	// 可以使用的CPU核数。
	Cpu *float64 `json:"cpu,omitempty"`

	// 可以使用的内存。
	Memory *string `json:"memory,omitempty"`
}

func (o ResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSpec struct{}"
	}

	return strings.Join([]string{"ResourceSpec", string(data)}, " ")
}
