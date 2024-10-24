package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagFabricWorkspace FabricWorkspace列表响应体。
type TagFabricWorkspace struct {

	// workspace的ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源详情。 资源对象，用于扩展。默认为空。
	ResourceDetail *interface{} `json:"resource_detail,omitempty"`

	// workspace名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 标签列表，没有标签默认为空数组。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	SysTags *[]SystemTag `json:"sys_tags,omitempty"`
}

func (o TagFabricWorkspace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagFabricWorkspace struct{}"
	}

	return strings.Join([]string{"TagFabricWorkspace", string(data)}, " ")
}
