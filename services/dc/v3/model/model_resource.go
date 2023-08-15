package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Resource 资源信息。
type Resource struct {

	// 资源详情。 资源对象，用于扩展。默认为空。
	ResourceDetail *interface{} `json:"resource_detail,omitempty"`

	// 资源的ID。
	ResourceId string `json:"resource_id"`

	// 资源名称，资源没有名称时默认为空字符串。
	ResourceName *string `json:"resource_name,omitempty"`

	// 标签列表，没有标签默认为空数组。
	Tags []Tag `json:"tags"`

	// 标签列表，没有标签默认为空数组。
	SysTags *[]Tag `json:"sys_tags,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
