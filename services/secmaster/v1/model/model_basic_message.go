package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BasicMessage **参数解释**: 错误描述 **取值范围**: 不涉及
type BasicMessage struct {
}

func (o BasicMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicMessage struct{}"
	}

	return strings.Join([]string{"BasicMessage", string(data)}, " ")
}
