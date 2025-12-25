package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Code **参数解释**: 错误码 **取值范围**: 不涉及
type Code struct {
}

func (o Code) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Code struct{}"
	}

	return strings.Join([]string{"Code", string(data)}, " ")
}
