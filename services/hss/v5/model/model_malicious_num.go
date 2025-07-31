package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MaliciousNum **参数解释**: 识别恶意进程数 **取值范围**: 最小值0，最大值2147483647
type MaliciousNum struct {
}

func (o MaliciousNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaliciousNum struct{}"
	}

	return strings.Join([]string{"MaliciousNum", string(data)}, " ")
}
