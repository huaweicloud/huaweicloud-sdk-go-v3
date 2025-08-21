package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportProfileVo struct {

	// **参数解释**： 模板ID **取值范围**： 不涉及
	ProfileId *string `json:"profile_id,omitempty"`

	// **参数解释**： 模板名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 模板类型 **取值范围**： daily 日报 weekly 周报 custom 自定义报告
	Category *string `json:"category,omitempty"`

	// **参数解释**： 启用状态 **取值范围**： 不涉及
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 最新的报告的ID **取值范围**： 不涉及
	ReportId *string `json:"report_id,omitempty"`

	// **参数解释**： 最新的报告的生成时间 **取值范围**： 不涉及
	LastTime *int64 `json:"last_time,omitempty"`
}

func (o ReportProfileVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportProfileVo struct{}"
	}

	return strings.Join([]string{"ReportProfileVo", string(data)}, " ")
}
