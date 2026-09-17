package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchSqlLimitingRuleNewResponse Response Object
type SwitchSqlLimitingRuleNewResponse struct {

	// 开关状态
	SwitchOn       *string `json:"switch_on,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchSqlLimitingRuleNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSqlLimitingRuleNewResponse struct{}"
	}

	return strings.Join([]string{"SwitchSqlLimitingRuleNewResponse", string(data)}, " ")
}
