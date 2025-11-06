package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmRuleTemplateResponse Response Object
type DeleteAlarmRuleTemplateResponse struct {

	// 告警模板状态列表。
	Resources      *[]DeleteAlarmRuleTemplateItemResult `json:"resources,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o DeleteAlarmRuleTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRuleTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRuleTemplateResponse", string(data)}, " ")
}
