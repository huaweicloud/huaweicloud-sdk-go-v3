package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchTransferRuleBody 批量转发规则
type BatchTransferRuleBody struct {

	// 批量转发规则
	Rules *[]TransferRuleBody `json:"rules,omitempty"`
}

func (o BatchTransferRuleBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTransferRuleBody struct{}"
	}

	return strings.Join([]string{"BatchTransferRuleBody", string(data)}, " ")
}
