package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrgConformancePackDetailedStatus 组织合规规则包查询结果。
type OrgConformancePackDetailedStatus struct {

	// 成员帐号ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 合规规则包ID。
	ConformancePackId *string `json:"conformance_pack_id,omitempty"`

	// 合规规则包名称。
	ConformancePackName *string `json:"conformance_pack_name,omitempty"`

	// 合规规则包部署状态
	State *string `json:"state,omitempty"`

	// 部署或删除组织合规规则包错误时的错误信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o OrgConformancePackDetailedStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrgConformancePackDetailedStatus struct{}"
	}

	return strings.Join([]string{"OrgConformancePackDetailedStatus", string(data)}, " ")
}
