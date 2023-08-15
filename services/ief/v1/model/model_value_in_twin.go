package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValueInTwin 终端设备动态属性
type ValueInTwin struct {
	Excepted *Excepted `json:"excepted,omitempty"`

	// 动态属性的实际信息
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
