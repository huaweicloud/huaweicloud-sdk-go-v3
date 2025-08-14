package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentIdRes **参数解释**: Agent ID **取值范围**: 字符长度1-64位
type AgentIdRes struct {
}

func (o AgentIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentIdRes struct{}"
	}

	return strings.Join([]string{"AgentIdRes", string(data)}, " ")
}
