package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessHandleMethod2 **参数解释**： 处理状态 **约束限制**: 不涉及 **取值范围**: - mark_as_trust：标记为可信 - mark_as_suspicious：标记为可疑 - isolate_and_kill：隔离查杀  **默认取值**: 不涉及
type ProcessHandleMethod2 struct {
}

func (o ProcessHandleMethod2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessHandleMethod2 struct{}"
	}

	return strings.Join([]string{"ProcessHandleMethod2", string(data)}, " ")
}
