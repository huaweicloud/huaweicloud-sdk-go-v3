package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MappingMessage **参数解释**: 错误描述 **取值范围**: 不涉及
type MappingMessage struct {
}

func (o MappingMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MappingMessage struct{}"
	}

	return strings.Join([]string{"MappingMessage", string(data)}, " ")
}
