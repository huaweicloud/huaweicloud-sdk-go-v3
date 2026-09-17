package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseSqlLimitRuleNewResponse Response Object
type ParseSqlLimitRuleNewResponse struct {

	// SQL限流规则
	Rule           *string `json:"rule,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ParseSqlLimitRuleNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseSqlLimitRuleNewResponse struct{}"
	}

	return strings.Join([]string{"ParseSqlLimitRuleNewResponse", string(data)}, " ")
}
