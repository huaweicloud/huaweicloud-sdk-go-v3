package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FilterTagResource 标签资源
type FilterTagResource struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceName string `json:"resource_name"`

	// 资源详情
	ResourceDetail *string `json:"resource_detail,omitempty"`

	// 资源下包含的标签
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o FilterTagResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterTagResource struct{}"
	}

	return strings.Join([]string{"FilterTagResource", string(data)}, " ")
}
