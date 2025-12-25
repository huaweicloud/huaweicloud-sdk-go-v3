package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListOneClickAlarmsRespOneClickAlarms struct {

	// **参数解释**： 一键告警ID。 **约束限制**： 不涉及。 **取值范围**： 只能为字母或者数字，字符长度为[1,64] **默认取值**： 不涉及。
	OneClickAlarmId string `json:"one_click_alarm_id"`

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及。
	Namespace string `json:"namespace"`

	// **参数解释** 一键告警描述 **约束限制** 不涉及 **取值范围** 字符长度在0到256之间 **默认取值** 空字符串
	Description string `json:"description"`

	// **参数解释** 是否启用一键告警 **约束限制** 不涉及 **取值范围** - true:开启 - false：关闭 **默认取值** true
	Enabled bool `json:"enabled"`
}

func (o ListOneClickAlarmsRespOneClickAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOneClickAlarmsRespOneClickAlarms struct{}"
	}

	return strings.Join([]string{"ListOneClickAlarmsRespOneClickAlarms", string(data)}, " ")
}
