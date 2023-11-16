package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceIpRuleResponse Response Object
type ListInstanceIpRuleResponse struct {

	// 转发规则总数
	Total *int32 `json:"total,omitempty"`

	// 转发规则列表
	Rules          *[]TransferRuleInfo `json:"rules,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListInstanceIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceIpRuleResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceIpRuleResponse", string(data)}, " ")
}
