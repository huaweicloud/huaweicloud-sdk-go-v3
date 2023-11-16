package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInstanceIpRuleResponse Response Object
type BatchDeleteInstanceIpRuleResponse struct {

	// 数量
	SuccessNum     *int32 `json:"success_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchDeleteInstanceIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceIpRuleResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceIpRuleResponse", string(data)}, " ")
}
