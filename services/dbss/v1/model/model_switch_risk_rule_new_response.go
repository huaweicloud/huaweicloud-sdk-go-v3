package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchRiskRuleNewResponse Response Object
type SwitchRiskRuleNewResponse struct {

	// 响应状态 - SUCCESS: 成功  - FAILED: 失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchRiskRuleNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchRiskRuleNewResponse struct{}"
	}

	return strings.Join([]string{"SwitchRiskRuleNewResponse", string(data)}, " ")
}
