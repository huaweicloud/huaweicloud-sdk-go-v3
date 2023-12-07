package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicResource Resource字段数据结构说明
type PublicResource struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源详情。资源对象，用于扩展。默认为空
	ResourceDetail *interface{} `json:"resource_detail"`

	// 标签列表，没有标签默认为空数组。请参考表ResourceTag字段数据结构说明。
	Tags []PublicResourceTag `json:"tags"`

	// 资源名称，没有默认为空字符串
	ResourceName string `json:"resource_name"`
}

func (o PublicResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicResource struct{}"
	}

	return strings.Join([]string{"PublicResource", string(data)}, " ")
}
