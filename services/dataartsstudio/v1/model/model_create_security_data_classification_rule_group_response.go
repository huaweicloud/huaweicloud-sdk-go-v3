package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityDataClassificationRuleGroupResponse Response Object
type CreateSecurityDataClassificationRuleGroupResponse struct {

	// 规则组ID
	Uuid *string `json:"uuid,omitempty"`

	// 规则组名称
	Name *string `json:"name,omitempty"`

	// 规则实体
	Rules *[]DataClassificationRuleQueryDto `json:"rules,omitempty"`

	// 规则组描述
	Description *string `json:"description,omitempty"`

	// 规则组创建人
	CreatedBy *string `json:"created_by,omitempty"`

	// 规则组创建时间
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 规则组更新人
	UpdatedBy *string `json:"updated_by,omitempty"`

	// 规则组更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// 项目ID
	ProjectId      *string `json:"project_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSecurityDataClassificationRuleGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityDataClassificationRuleGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityDataClassificationRuleGroupResponse", string(data)}, " ")
}
