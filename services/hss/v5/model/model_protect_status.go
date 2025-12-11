package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectStatus **参数解释**: 防护状态 **取值范围**:  - closed ：关闭  - opened ：开启
type ProtectStatus struct {
}

func (o ProtectStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectStatus struct{}"
	}

	return strings.Join([]string{"ProtectStatus", string(data)}, " ")
}
