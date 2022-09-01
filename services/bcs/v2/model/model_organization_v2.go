package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 组织信息
type OrganizationV2 struct {

	// 组织名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 组织hash
	NameHash *string `json:"name_hash,omitempty" xml:"name_hash"`

	// 组织节点
	NodeCount *string `json:"node_count,omitempty" xml:"node_count"`

	// 状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 状态描述
	StatusDetail *string `json:"status_detail,omitempty" xml:"status_detail"`
}

func (o OrganizationV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationV2 struct{}"
	}

	return strings.Join([]string{"OrganizationV2", string(data)}, " ")
}
