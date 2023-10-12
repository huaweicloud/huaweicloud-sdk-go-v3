package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrgConformancePackStatus 组织合规规则包详情。
type OrgConformancePackStatus struct {

	// 组织合规规则包ID。
	OrgConformancePackId *string `json:"org_conformance_pack_id,omitempty"`

	// 组织合规规则包名称。
	OrgConformancePackName *string `json:"org_conformance_pack_name,omitempty"`

	// 合规规则包部署状态
	State *string `json:"state,omitempty"`

	// 部署或删除组织合规规则包错误时的错误信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 组织合规规则包创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 组织合规规则包更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o OrgConformancePackStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrgConformancePackStatus struct{}"
	}

	return strings.Join([]string{"OrgConformancePackStatus", string(data)}, " ")
}
