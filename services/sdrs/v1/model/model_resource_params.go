package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源参数数据结构
type ResourceParams struct {

	// 保护实例ID。
	ResourceId string `json:"resource_id"`

	ResourceDetail *ShowProtectedInstanceParams `json:"resource_detail"`

	// 标签列表，没有标签默认为空数组。
	Tags []ResourceTag `json:"tags"`

	// 保护实例名称，没有名称时默认为空字符串。
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ResourceParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceParams struct{}"
	}

	return strings.Join([]string{"ResourceParams", string(data)}, " ")
}
