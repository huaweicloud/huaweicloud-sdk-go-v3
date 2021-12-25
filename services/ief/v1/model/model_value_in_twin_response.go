package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 终端设备静态属性信息
type ValueInTwinResponse struct {
	Excepted *ValueInTwinResponseExcepted `json:"excepted,omitempty"`

	Actual *ValueInTwinResponseActual `json:"actual,omitempty"`

	Metadata *ValueInTwinResponseMetadata `json:"metadata,omitempty"`
	// 标识属性是否可选，默认为true，继承模板的属性默认为false

	Optional *bool `json:"optional,omitempty"`
}

func (o ValueInTwinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueInTwinResponse struct{}"
	}

	return strings.Join([]string{"ValueInTwinResponse", string(data)}, " ")
}
