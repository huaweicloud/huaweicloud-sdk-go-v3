package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleRespData **参数解释**： 时间表返回体 **取值范围**： 不涉及
type ScheduleRespData struct {

	// **参数解释**： 时间表ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o ScheduleRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleRespData struct{}"
	}

	return strings.Join([]string{"ScheduleRespData", string(data)}, " ")
}
