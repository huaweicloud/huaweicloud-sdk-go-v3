package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectionEnabledTypes 已开启防护类型
type ProtectionEnabledTypes struct {

	// **参数解释**: 防护类型 **取值范围**: - host：主机 - container：容器 - application：应用 - serverless：Serverless
	ProtectType *string `json:"protect_type,omitempty"`
}

func (o ProtectionEnabledTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectionEnabledTypes struct{}"
	}

	return strings.Join([]string{"ProtectionEnabledTypes", string(data)}, " ")
}
