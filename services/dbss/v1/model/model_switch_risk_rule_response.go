package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SwitchRiskRuleResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchRiskRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchRiskRuleResponse struct{}"
	}

	return strings.Join([]string{"SwitchRiskRuleResponse", string(data)}, " ")
}
