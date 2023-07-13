package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeResource struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`
}

func (o NodeResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeResource struct{}"
	}

	return strings.Join([]string{"NodeResource", string(data)}, " ")
}
