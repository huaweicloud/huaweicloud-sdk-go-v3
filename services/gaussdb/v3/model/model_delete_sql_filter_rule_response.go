package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSqlFilterRuleResponse Response Object
type DeleteSqlFilterRuleResponse struct {

	// 删除SQL限流规则任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSqlFilterRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlFilterRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteSqlFilterRuleResponse", string(data)}, " ")
}
