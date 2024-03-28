package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTags 资源项。
type ResourceTags struct {

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o ResourceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTags struct{}"
	}

	return strings.Join([]string{"ResourceTags", string(data)}, " ")
}
