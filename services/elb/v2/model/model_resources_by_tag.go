package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 根据标签查询资源响应返回对象。
type ResourcesByTag struct {
	// 资源ID.

	ResourceId string `json:"resource_id"`
	// 资源名称。

	ResourceName string `json:"resource_name"`
	// 资源描述。

	ResourceDetail string `json:"resource_detail"`
	// 资源标签。

	Tags []ResourceTag `json:"tags"`
	// 父级资源ID。

	SuperResourceId *string `json:"super_resource_id,omitempty"`
}

func (o ResourcesByTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesByTag struct{}"
	}

	return strings.Join([]string{"ResourcesByTag", string(data)}, " ")
}
