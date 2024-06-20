package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseSqlLimitRulesResponse Response Object
type ParseSqlLimitRulesResponse struct {

	// SQL限流关键字
	Rule           *string `json:"rule,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ParseSqlLimitRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseSqlLimitRulesResponse struct{}"
	}

	return strings.Join([]string{"ParseSqlLimitRulesResponse", string(data)}, " ")
}
