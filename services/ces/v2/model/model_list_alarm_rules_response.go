package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmRulesResponse Response Object
type ListAlarmRulesResponse struct {

	// **参数解释**： 告警规则列表。 **取值范围**： 最多包含100个告警规则信息。
	Alarms *[]ListAlarmRespBodyAlarms `json:"alarms,omitempty"`

	// **参数解释**： 告警规则总数。 **取值范围**： [0,10000]
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmRulesResponse", string(data)}, " ")
}
