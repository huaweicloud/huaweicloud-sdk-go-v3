package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetSqlFilterRuleRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *OperateSqlFilterRuleReq `json:"body,omitempty"`
}

func (o SetSqlFilterRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSqlFilterRuleRequest struct{}"
	}

	return strings.Join([]string{"SetSqlFilterRuleRequest", string(data)}, " ")
}
