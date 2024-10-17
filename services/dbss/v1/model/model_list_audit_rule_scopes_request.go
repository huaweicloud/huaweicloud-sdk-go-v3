package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditRuleScopesRequest Request Object
type ListAuditRuleScopesRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 查询记录数
	Limit *string `json:"limit,omitempty"`
}

func (o ListAuditRuleScopesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditRuleScopesRequest struct{}"
	}

	return strings.Join([]string{"ListAuditRuleScopesRequest", string(data)}, " ")
}
