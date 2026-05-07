package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentIdRes **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **取值范围**: 字符长度1-64位
type AgentIdRes struct {
}

func (o AgentIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentIdRes struct{}"
	}

	return strings.Join([]string{"AgentIdRes", string(data)}, " ")
}
