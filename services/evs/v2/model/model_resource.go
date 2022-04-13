package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {
	// 资源ID。

	ResourceId string `json:"resource_id"`
	// 资源名称。

	ResourceName *string `json:"resource_name,omitempty"`

	ResourceDetail *VolumeDetailForTag `json:"resource_detail"`
	// 标签列表。

	Tags []map[string]string `json:"tags"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
