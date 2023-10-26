package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePermRuleRequest Request Object
type DeletePermRuleRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	// 权限规则id
	RuleId string `json:"rule_id"`
}

func (o DeletePermRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePermRuleRequest struct{}"
	}

	return strings.Join([]string{"DeletePermRuleRequest", string(data)}, " ")
}
