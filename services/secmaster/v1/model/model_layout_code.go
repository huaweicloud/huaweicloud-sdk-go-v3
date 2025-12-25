package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LayoutCode **参数解释**: 错误码 **取值范围**: 不涉及
type LayoutCode struct {
}

func (o LayoutCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayoutCode struct{}"
	}

	return strings.Join([]string{"LayoutCode", string(data)}, " ")
}
