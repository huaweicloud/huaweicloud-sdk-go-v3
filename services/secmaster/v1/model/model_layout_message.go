package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LayoutMessage **参数解释**: 错误描述 **取值范围**: 不涉及
type LayoutMessage struct {
}

func (o LayoutMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayoutMessage struct{}"
	}

	return strings.Join([]string{"LayoutMessage", string(data)}, " ")
}
