package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateInstanceIpRuleResponse Response Object
type BatchCreateInstanceIpRuleResponse struct {

	// 数量
	SuccessNum     *int32 `json:"success_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCreateInstanceIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInstanceIpRuleResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateInstanceIpRuleResponse", string(data)}, " ")
}
