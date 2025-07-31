package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LearningStatus **参数解释**： 服务器名称 **约束限制**: 不涉及 **取值范围**: - effecting：学习完成，策略生效 - learned：学习完成，待确认 - learning：学习中 - pause：暂停 - abnormal：学习异常  **默认取值**: 不涉及
type LearningStatus struct {
}

func (o LearningStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LearningStatus struct{}"
	}

	return strings.Join([]string{"LearningStatus", string(data)}, " ")
}
