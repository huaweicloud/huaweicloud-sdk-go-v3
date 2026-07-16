package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Value 资源量。
type Value struct {

	// cpu量，即计算资源量。
	Cpu *string `json:"cpu,omitempty"`

	// 内存。
	Memory *string `json:"memory,omitempty"`

	// GPU卡的数量。
	Tnt004 *string `json:"tnt004,omitempty"`
}

func (o Value) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Value struct{}"
	}

	return strings.Join([]string{"Value", string(data)}, " ")
}
