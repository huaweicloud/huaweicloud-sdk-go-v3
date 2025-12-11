package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FailReason **参数解释**: 失败原因 **取值范围**: 字符长度0-512位
type FailReason struct {
}

func (o FailReason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailReason struct{}"
	}

	return strings.Join([]string{"FailReason", string(data)}, " ")
}
