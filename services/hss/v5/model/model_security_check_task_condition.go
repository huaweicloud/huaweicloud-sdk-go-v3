package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckTaskCondition 基线检测周期扫描配置
type SecurityCheckTaskCondition struct {

	// **参数解释**： 定时任务类型 **取值范围**： - fixed_weekday : 固定工作日
	Type *string `json:"type,omitempty"`

	// 周几触发，可选0或多个
	DayOfWeek *[]string `json:"day_of_week,omitempty"`

	// **参数解释**： 在此参数表示的小时触发定时任务 **取值范围**： 最小值0，最大值23
	Hour *int32 `json:"hour,omitempty"`

	// **参数解释**： 在此参数表示的分钟触发定时任务 **取值范围**： 最小值0，最大值59
	Minute *int32 `json:"minute,omitempty"`

	// **参数解释**： 随机偏移时间 **取值范围**： 最小值0，最大值7200
	RandomOffset *int32 `json:"random_offset,omitempty"`
}

func (o SecurityCheckTaskCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckTaskCondition struct{}"
	}

	return strings.Join([]string{"SecurityCheckTaskCondition", string(data)}, " ")
}
