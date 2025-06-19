package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlJobDefendRuleRequest Request Object
type CreateSqlJobDefendRuleRequest struct {
	Body *SqlJobDefendRuleRequestBody `json:"body,omitempty"`
}

func (o CreateSqlJobDefendRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlJobDefendRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlJobDefendRuleRequest", string(data)}, " ")
}
