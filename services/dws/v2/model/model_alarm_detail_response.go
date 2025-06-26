package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmDetailResponse **参数解释**： 告警详情。 **取值范围**： 不涉及。
type AlarmDetailResponse struct {

	// **参数解释**： 告警定义ID。 **取值范围**： 不涉及。
	AlarmId *string `json:"alarm_id,omitempty"`

	// **参数解释**： 告警名称。 **取值范围**： 不涉及。
	AlarmName *string `json:"alarm_name,omitempty"`

	// **参数解释**： 告警级别。 **取值范围**： 不涉及。
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// **参数解释**： 告警服务。 **取值范围**： 不涉及。
	AlarmSource *string `json:"alarm_source,omitempty"`

	// **参数解释**： 告警消息。 **取值范围**： 不涉及。
	AlarmMessage *string `json:"alarm_message,omitempty"`

	// **参数解释**： 告警定位信息。 **取值范围**： 不涉及。
	AlarmLocation *string `json:"alarm_location,omitempty"`

	// **参数解释**： 告警源ID。 **取值范围**： 不涉及。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 告警源名称。 **取值范围**： 不涉及。
	ResourceIdName *string `json:"resource_id_name,omitempty"`

	// **参数解释**： 告警日期。 **取值范围**： 不涉及。
	AlarmGenerateDate *string `json:"alarm_generate_date,omitempty"`

	// **参数解释**： 告警状态。 **取值范围**： 不涉及。
	AlarmStatus *string `json:"alarm_status,omitempty"`
}

func (o AlarmDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmDetailResponse struct{}"
	}

	return strings.Join([]string{"AlarmDetailResponse", string(data)}, " ")
}
