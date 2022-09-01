package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceRelation struct {

	// 关系类型
	RelationType *string `json:"relation_type,omitempty" xml:"relation_type"`

	// 源资源类型
	FromResourceType *string `json:"from_resource_type,omitempty" xml:"from_resource_type"`

	// 目的资源类型
	ToResourceType *string `json:"to_resource_type,omitempty" xml:"to_resource_type"`

	// 源资源ID
	FromResourceId *string `json:"from_resource_id,omitempty" xml:"from_resource_id"`

	// 目的资源ID
	ToResourceId *string `json:"to_resource_id,omitempty" xml:"to_resource_id"`
}

func (o ResourceRelation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceRelation struct{}"
	}

	return strings.Join([]string{"ResourceRelation", string(data)}, " ")
}
