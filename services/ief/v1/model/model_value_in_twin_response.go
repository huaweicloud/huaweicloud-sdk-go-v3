package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 终端设备静态属性信息
type ValueInTwinResponse struct {
	Excepted *ExceptedActual `json:"excepted,omitempty" xml:"excepted"`

	Actual *ExceptedActual `json:"actual,omitempty" xml:"actual"`

	Metadata *Metadata `json:"metadata,omitempty" xml:"metadata"`

	// 标识属性是否可选，默认为true，继承模板的属性默认为false
	Optional *bool `json:"optional,omitempty" xml:"optional"`
}

func (o ValueInTwinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueInTwinResponse struct{}"
	}

	return strings.Join([]string{"ValueInTwinResponse", string(data)}, " ")
}
