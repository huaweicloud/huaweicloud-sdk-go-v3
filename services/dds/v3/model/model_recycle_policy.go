package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecyclePolicy struct {

	// **参数解释：** 打开回收策略。 **约束限制：** 不可关闭。 **取值范围：** true。 **默认取值：** true。
	Enabled bool `json:"enabled"`

	// **参数解释：** 策略保持时长（1-7天），天数为正整数，不填默认保留7天。 **约束限制：** 正整数。 **取值范围：** 1-7。 **默认取值：** 7。
	RetentionPeriodInDays *int32 `json:"retention_period_in_days,omitempty"`
}

func (o RecyclePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecyclePolicy struct{}"
	}

	return strings.Join([]string{"RecyclePolicy", string(data)}, " ")
}
