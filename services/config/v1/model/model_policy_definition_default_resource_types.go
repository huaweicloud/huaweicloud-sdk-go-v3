package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyDefinitionDefaultResourceTypes 默认资源类型
type PolicyDefinitionDefaultResourceTypes struct {

	// 云服务名称
	Provider *string `json:"provider,omitempty"`

	// 资源类型
	Type *string `json:"type,omitempty"`
}

func (o PolicyDefinitionDefaultResourceTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyDefinitionDefaultResourceTypes struct{}"
	}

	return strings.Join([]string{"PolicyDefinitionDefaultResourceTypes", string(data)}, " ")
}
