package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateKeywordsAlarmRuleRequest struct {
	Body *UpdateKeywordsAlarmRuleRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateKeywordsAlarmRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeywordsAlarmRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeywordsAlarmRuleRequest", string(data)}, " ")
}
