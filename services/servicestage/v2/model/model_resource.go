package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {

	// 资源ID
	Id string `json:"id"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	Type *ResourceType `json:"type"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
