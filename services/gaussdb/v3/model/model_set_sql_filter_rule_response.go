package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SetSqlFilterRuleResponse struct {

	// 设置SQL限流规则任务ID。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o SetSqlFilterRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSqlFilterRuleResponse struct{}"
	}

	return strings.Join([]string{"SetSqlFilterRuleResponse", string(data)}, " ")
}
