package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 终端设备动态属性
type ValueInTwin struct {
	Excepted *ValueInTwinExcepted `json:"excepted"`
	// 动态属性的期望信息

	Optional *bool `json:"optional,omitempty"`

	Metadata *Metadata `json:"metadata,omitempty"`
}

func (o ValueInTwin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueInTwin struct{}"
	}

	return strings.Join([]string{"ValueInTwin", string(data)}, " ")
}
