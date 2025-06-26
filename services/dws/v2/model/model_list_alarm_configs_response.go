package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmConfigsResponse Response Object
type ListAlarmConfigsResponse struct {

	// **参数解释**： 告警配置总数。 **取值范围**： 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 告警配置列表。 **取值范围**： 不涉及。
	AlarmConfigs   *[]AlarmConfigResponse `json:"alarm_configs,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAlarmConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmConfigsResponse", string(data)}, " ")
}
