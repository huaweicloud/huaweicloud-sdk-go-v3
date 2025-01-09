package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSharesByTagResource 资源详情
type ListSharesByTagResource struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源详情
	ResourceDetail *string `json:"resource_detail,omitempty"`

	// 资源的标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o ListSharesByTagResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharesByTagResource struct{}"
	}

	return strings.Join([]string{"ListSharesByTagResource", string(data)}, " ")
}
