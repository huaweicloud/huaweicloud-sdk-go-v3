package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRuleFixFailDetailInfo 配置检测检查项修复失败原因
type CheckRuleFixFailDetailInfo struct {

	// **参数解释**: 修复失败原因 **取值范围**: 字符长度0-65534位
	FixFailReason *string `json:"fix_fail_reason,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`
}

func (o CheckRuleFixFailDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRuleFixFailDetailInfo struct{}"
	}

	return strings.Join([]string{"CheckRuleFixFailDetailInfo", string(data)}, " ")
}
