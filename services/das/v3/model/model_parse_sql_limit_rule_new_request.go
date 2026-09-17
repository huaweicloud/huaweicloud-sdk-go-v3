package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseSqlLimitRuleNewRequest Request Object
type ParseSqlLimitRuleNewRequest struct {
	Body *ParseSqlLimitRuleNewRequestBody `json:"body,omitempty"`
}

func (o ParseSqlLimitRuleNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseSqlLimitRuleNewRequest struct{}"
	}

	return strings.Join([]string{"ParseSqlLimitRuleNewRequest", string(data)}, " ")
}
