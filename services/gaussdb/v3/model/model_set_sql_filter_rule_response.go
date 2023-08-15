package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSqlFilterRuleResponse Response Object
type SetSqlFilterRuleResponse struct {

	// 设置SQL限流规则任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetSqlFilterRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSqlFilterRuleResponse struct{}"
	}

	return strings.Join([]string{"SetSqlFilterRuleResponse", string(data)}, " ")
}
