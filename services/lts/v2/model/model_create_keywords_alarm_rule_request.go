package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateKeywordsAlarmRuleRequest struct {
	Body *CreateKeywordsAlarmRuleRequestBody `json:"body,omitempty"`
}

func (o CreateKeywordsAlarmRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeywordsAlarmRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateKeywordsAlarmRuleRequest", string(data)}, " ")
}
