package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmTemplateAssociationAlarmsResponse Response Object
type ListAlarmTemplateAssociationAlarmsResponse struct {

	// 告警规则列表
	Alarms *[]ListAssociationAlarmsResponseAlarms `json:"alarms,omitempty"`

	// **参数解释**： 告警规则列表总数。     **取值范围**： 取值范围为0-1000.
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmTemplateAssociationAlarmsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTemplateAssociationAlarmsResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmTemplateAssociationAlarmsResponse", string(data)}, " ")
}
