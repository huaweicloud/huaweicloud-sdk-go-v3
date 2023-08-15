package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlInjectionRulesRequest Request Object
type ListSqlInjectionRulesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SqlRuleRequest `json:"body,omitempty"`
}

func (o ListSqlInjectionRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlInjectionRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlInjectionRulesRequest", string(data)}, " ")
}
