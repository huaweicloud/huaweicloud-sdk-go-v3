package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableM **参数解释**: 启用、停用 **约束限制**: 必填 **取值范围**: - 1：启用 - 0：停用  **默认取值**: 不涉及
type EnableM struct {
}

func (o EnableM) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableM struct{}"
	}

	return strings.Join([]string{"EnableM", string(data)}, " ")
}
