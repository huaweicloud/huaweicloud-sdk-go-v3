package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源实例详情
type ResourceInstance struct {
	// 资源ID，Endpoint Service ID或Endpoint ID。

	ResourceId *string `json:"resource_id,omitempty"`
	// 标签列表。

	Tags *[]TagList `json:"tags,omitempty"`
	// 资源名称，资源没有名称时，返回ID。

	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ResourceInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstance struct{}"
	}

	return strings.Join([]string{"ResourceInstance", string(data)}, " ")
}
