package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePermRuleRequest Request Object
type UpdatePermRuleRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	// 权限规则id
	RuleId string `json:"rule_id"`

	Body *OnePermRuleRequestInfo `json:"body,omitempty"`
}

func (o UpdatePermRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdatePermRuleRequest", string(data)}, " ")
}
