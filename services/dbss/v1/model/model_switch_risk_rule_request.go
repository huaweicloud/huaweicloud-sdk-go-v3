package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchRiskRuleRequest Request Object
type SwitchRiskRuleRequest struct {

	// 实例ID。可在查询实例列表接口的ID字段获取。
	InstanceId string `json:"instance_id"`

	Body *BatchSwitchesRequest `json:"body,omitempty"`
}

func (o SwitchRiskRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchRiskRuleRequest struct{}"
	}

	return strings.Join([]string{"SwitchRiskRuleRequest", string(data)}, " ")
}
