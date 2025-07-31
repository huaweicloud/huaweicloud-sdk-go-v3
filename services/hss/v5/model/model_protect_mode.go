package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectMode 防护模式，包含如下4种   - recovery ：恢复模式   - alarm ：告警模式
type ProtectMode struct {
}

func (o ProtectMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectMode struct{}"
	}

	return strings.Join([]string{"ProtectMode", string(data)}, " ")
}
