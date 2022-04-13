package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SwitchRiskRuleRequest struct {
	// instance_id

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
