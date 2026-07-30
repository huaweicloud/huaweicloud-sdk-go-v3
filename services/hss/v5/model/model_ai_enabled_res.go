package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiEnabledRes **参数解释**: 是否启用 **取值范围**: - false：否 - true：是
type AiEnabledRes struct {
}

func (o AiEnabledRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiEnabledRes struct{}"
	}

	return strings.Join([]string{"AiEnabledRes", string(data)}, " ")
}
