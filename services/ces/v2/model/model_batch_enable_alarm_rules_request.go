package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableAlarmRulesRequest Request Object
type BatchEnableAlarmRulesRequest struct {
	Body *BatchEnableAlarmsRequestBody `json:"body,omitempty"`
}

func (o BatchEnableAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchEnableAlarmRulesRequest", string(data)}, " ")
}
