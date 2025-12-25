package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmTypeResp **参数解释**： 告警类型。 **取值范围**： 针对事件类型的告警时，告警类型为EVENT.SYS（系统事件）或EVENT.CUSTOM（自定义事件）。 针对资源分组的告警时，告警类型为RESOURCE_GROUP。 针对指定资源的告警时，告警类型为MULTI_INSTANCE。 - EVENT.SYS：针对系统事件的告警类型 - EVENT.CUSTOM：针对自定义事件的告警类型 - RESOURCE_GROUP：针对资源分组的告警类型 - MULTI_INSTANCE：指定资源的告警类型
type AlarmTypeResp struct {
}

func (o AlarmTypeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTypeResp struct{}"
	}

	return strings.Join([]string{"AlarmTypeResp", string(data)}, " ")
}
