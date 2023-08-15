package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceDto 资源信息。
type ResourceDto struct {

	// 资源Id。
	ResourceId string `json:"resource_id"`

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源标签列表。
	Tags []Match `json:"tags"`
}

func (o ResourceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDto struct{}"
	}

	return strings.Join([]string{"ResourceDto", string(data)}, " ")
}
