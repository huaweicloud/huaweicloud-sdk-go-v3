package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MappingCode **参数解释**: 错误码 **取值范围**: 不涉及
type MappingCode struct {
}

func (o MappingCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MappingCode struct{}"
	}

	return strings.Join([]string{"MappingCode", string(data)}, " ")
}
