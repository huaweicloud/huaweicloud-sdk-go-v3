package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiIsExclusiveRes **参数解释**: 是否是默认策略 **取值范围**: - false：否 - true：是
type AiIsExclusiveRes struct {
}

func (o AiIsExclusiveRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiIsExclusiveRes struct{}"
	}

	return strings.Join([]string{"AiIsExclusiveRes", string(data)}, " ")
}
