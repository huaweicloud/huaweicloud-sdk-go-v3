package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlInjectionRulesRequest Request Object
type ListSqlInjectionRulesRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
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
