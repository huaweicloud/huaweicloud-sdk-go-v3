package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmTemplateAssociationAlarmsRequest Request Object
type ListAlarmTemplateAssociationAlarmsRequest struct {

	// **参数解释**： 分页偏移量 **约束限制**： 不涉及 **取值范围**： 整数，[0,10000] **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 分页大小 **约束限制**： 不涉及 **取值范围**： 整数，[1,100] **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 告警模板的ID。     **约束限制**： 不涉及。 **取值范围**： 以at开头，后跟字母、数字，长度为[2,64]个字符。           **默认取值**： 不涉及。
	TemplateId string `json:"template_id"`
}

func (o ListAlarmTemplateAssociationAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTemplateAssociationAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmTemplateAssociationAlarmsRequest", string(data)}, " ")
}
