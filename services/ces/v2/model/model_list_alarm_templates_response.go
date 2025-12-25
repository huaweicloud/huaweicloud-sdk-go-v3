package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmTemplatesResponse Response Object
type ListAlarmTemplatesResponse struct {

	// 告警模板列表
	AlarmTemplates *[]AlarmTemplates `json:"alarm_templates,omitempty"`

	// **参数解释**： 告警模板记录总数。 **取值范围**： 字符串长度在 10 到 9999999 之间。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmTemplatesResponse", string(data)}, " ")
}
