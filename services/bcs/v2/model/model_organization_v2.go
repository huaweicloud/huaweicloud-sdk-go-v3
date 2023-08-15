package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationV2 组织信息
type OrganizationV2 struct {

	// 组织名称
	Name *string `json:"name,omitempty"`

	// 组织hash
	NameHash *string `json:"name_hash,omitempty"`

	// 组织节点
	NodeCount *int32 `json:"node_count,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 状态描述
	StatusDetail *string `json:"status_detail,omitempty"`
}

func (o OrganizationV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationV2 struct{}"
	}

	return strings.Join([]string{"OrganizationV2", string(data)}, " ")
}
