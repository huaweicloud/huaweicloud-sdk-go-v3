package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Message **参数解释**: 错误描述 **取值范围**: 不涉及
type Message struct {
}

func (o Message) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Message struct{}"
	}

	return strings.Join([]string{"Message", string(data)}, " ")
}
