package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorAttribute 规格属性信息
type FlavorAttribute struct {

	// 规格属性名称，如mem、cpu。
	Name *string `json:"name,omitempty"`

	// 规格属性值
	Value *string `json:"value,omitempty"`
}

func (o FlavorAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorAttribute struct{}"
	}

	return strings.Join([]string{"FlavorAttribute", string(data)}, " ")
}
