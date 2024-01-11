package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceInstance 查询标签返回资源信息体
type ResourceInstance struct {

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源详情。 资源对象，用于扩展。默认为空。
	ResourceDetail *string `json:"resource_detail,omitempty"`

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 标签列表。
	Tags *[]ResourceTagInfo `json:"tags,omitempty"`
}

func (o ResourceInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstance struct{}"
	}

	return strings.Join([]string{"ResourceInstance", string(data)}, " ")
}
