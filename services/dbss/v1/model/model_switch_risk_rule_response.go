package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchRiskRuleResponse Response Object
type SwitchRiskRuleResponse struct {

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchRiskRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchRiskRuleResponse struct{}"
	}

	return strings.Join([]string{"SwitchRiskRuleResponse", string(data)}, " ")
}
