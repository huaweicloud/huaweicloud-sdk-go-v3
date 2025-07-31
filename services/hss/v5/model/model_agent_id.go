package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentId **参数解释**: Agent ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
type AgentId struct {
}

func (o AgentId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentId struct{}"
	}

	return strings.Join([]string{"AgentId", string(data)}, " ")
}
