package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAssociationAlarmsResponseAlarms struct {

	// **参数解释**： 告警模板的ID。  **取值范围**：    以al开头，后跟22位由字母或数字组成的字符串。长度为24个字符。
	AlarmId string `json:"alarm_id"`

	// **参数解释**： 告警规则名称。     **取值范围**： 只能为字母、数字、汉字、-或_，长度为[1,128]个字符
	Name string `json:"name"`

	// **参数解释**： 告警规则描述。     **取值范围**： 长度为[0,256]个字符。
	Description string `json:"description"`
}

func (o ListAssociationAlarmsResponseAlarms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociationAlarmsResponseAlarms struct{}"
	}

	return strings.Join([]string{"ListAssociationAlarmsResponseAlarms", string(data)}, " ")
}
