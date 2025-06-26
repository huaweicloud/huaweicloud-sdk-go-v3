package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmConfigResponse **参数解释**： 告警配置详情。 **取值范围**： 不涉及。
type AlarmConfigResponse struct {

	// **参数解释**： 告警配置ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 告警ID。 **取值范围**： 不涉及。
	AlarmId *string `json:"alarm_id,omitempty"`

	// **参数解释**： 告警名称。 **取值范围**： 不涉及。
	AlarmName *string `json:"alarm_name,omitempty"`

	// **参数解释**： 所属服务。 **取值范围**： 不涉及。
	NameSpace *string `json:"name_space,omitempty"`

	// **参数解释**： 告警级别。 **取值范围**： 不涉及。
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// **参数解释**： 用户是否可见。 **取值范围**： 不涉及。
	IsUserVisible *string `json:"is_user_visible,omitempty"`

	// **参数解释**： 是否覆盖。 **取值范围**： 不涉及。
	IsConverge *string `json:"is_converge,omitempty"`

	// **参数解释**： 覆盖时间。 **取值范围**： 不涉及。
	ConvergeTime *int32 `json:"converge_time,omitempty"`

	// **参数解释**： 运维是否可见。 **取值范围**： 不涉及。
	IsMaintainVisible *string `json:"is_maintain_visible,omitempty"`
}

func (o AlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"AlarmConfigResponse", string(data)}, " ")
}
