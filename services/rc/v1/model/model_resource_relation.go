package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceRelation struct {

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 关联资源类型
	RelatedResourceType *string `json:"related_resource_type,omitempty"`

	// 关联资源ID
	RelatedResourceId *string `json:"related_resource_id,omitempty"`
}

func (o ResourceRelation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceRelation struct{}"
	}

	return strings.Join([]string{"ResourceRelation", string(data)}, " ")
}
