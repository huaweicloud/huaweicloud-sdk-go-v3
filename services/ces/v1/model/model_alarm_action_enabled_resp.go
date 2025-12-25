package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmActionEnabledResp **参数解释**： 该条告警触发时，是否开启告警通知。 **取值范围**： 布尔值。 - true：开启告警通知。 - false：不开启告警通知。
type AlarmActionEnabledResp struct {
}

func (o AlarmActionEnabledResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmActionEnabledResp struct{}"
	}

	return strings.Join([]string{"AlarmActionEnabledResp", string(data)}, " ")
}
