package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmEnabledResp **参数解释**： 是否启用该条告警。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。
type AlarmEnabledResp struct {
}

func (o AlarmEnabledResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmEnabledResp struct{}"
	}

	return strings.Join([]string{"AlarmEnabledResp", string(data)}, " ")
}
