package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SuspiciousNum **参数解释**: 识别可疑进程数 **取值范围**: 最小值0，最大值2147483647
type SuspiciousNum struct {
}

func (o SuspiciousNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SuspiciousNum struct{}"
	}

	return strings.Join([]string{"SuspiciousNum", string(data)}, " ")
}
