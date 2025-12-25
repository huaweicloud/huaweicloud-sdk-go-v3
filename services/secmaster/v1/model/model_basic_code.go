package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BasicCode **参数解释**: 错误码 **取值范围**: 不涉及
type BasicCode struct {
}

func (o BasicCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicCode struct{}"
	}

	return strings.Join([]string{"BasicCode", string(data)}, " ")
}
