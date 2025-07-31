package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LearningDays **参数解释**: 策略学习天数 **取值范围**: 最小值1，最大值1000
type LearningDays struct {
}

func (o LearningDays) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LearningDays struct{}"
	}

	return strings.Join([]string{"LearningDays", string(data)}, " ")
}
