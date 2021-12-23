package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {
	// 资源id

	ResourceId *string `json:"resource_id,omitempty"`
	// 资源类型

	ResourceType *string `json:"resource_type,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
