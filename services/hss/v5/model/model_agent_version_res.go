package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentVersionRes **参数解释**： agent版本 **取值范围**： 字符长度1-32位
type AgentVersionRes struct {
}

func (o AgentVersionRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentVersionRes struct{}"
	}

	return strings.Join([]string{"AgentVersionRes", string(data)}, " ")
}
