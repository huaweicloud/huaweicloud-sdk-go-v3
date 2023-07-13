package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAlarmRulesRequest Request Object
type BatchDeleteAlarmRulesRequest struct {
	Body *BatchDeleteAlarmsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAlarmRulesRequest", string(data)}, " ")
}
