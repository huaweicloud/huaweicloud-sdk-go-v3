package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Action **参数解释**: 处置动作 **取值范围**: - auto：自动处置 - manual：人工处置
type Action struct {
}

func (o Action) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Action struct{}"
	}

	return strings.Join([]string{"Action", string(data)}, " ")
}
