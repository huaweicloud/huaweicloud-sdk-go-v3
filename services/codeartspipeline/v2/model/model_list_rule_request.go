package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuleRequest Request Object
type ListRuleRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 项目ID
	CloudProjectId *string `json:"cloud_project_id,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0
	Offset int32 `json:"offset"`

	// 每页显示的条目数量
	Limit int32 `json:"limit"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 规则名称，用于模糊搜索
	Name *string `json:"name,omitempty"`
}

func (o ListRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleRequest struct{}"
	}

	return strings.Join([]string{"ListRuleRequest", string(data)}, " ")
}
