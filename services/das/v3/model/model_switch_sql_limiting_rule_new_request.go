package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchSqlLimitingRuleNewRequest Request Object
type SwitchSqlLimitingRuleNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SwitchSqlLimitingRuleNewRequestBody `json:"body,omitempty"`
}

func (o SwitchSqlLimitingRuleNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSqlLimitingRuleNewRequest struct{}"
	}

	return strings.Join([]string{"SwitchSqlLimitingRuleNewRequest", string(data)}, " ")
}
