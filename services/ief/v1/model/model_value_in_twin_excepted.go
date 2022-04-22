package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 动态属性的期望信息
type ValueInTwinExcepted struct {

	// 动态属性的初始值，最大长度512， value允许英文字母、数字、下划线、中划线、点、逗号、冒号、/、@、#
	Value string `json:"value"`

	Metadata *ValueInTwinExceptedMetadata `json:"metadata,omitempty"`
}

func (o ValueInTwinExcepted) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueInTwinExcepted struct{}"
	}

	return strings.Join([]string{"ValueInTwinExcepted", string(data)}, " ")
}
