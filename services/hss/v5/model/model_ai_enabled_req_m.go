package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiEnabledReqM **参数解释**: 是否启用 **约束限制**: 必填 **取值范围**: - false：否 - true：是  **默认取值**: 不涉及
type AiEnabledReqM struct {
}

func (o AiEnabledReqM) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiEnabledReqM struct{}"
	}

	return strings.Join([]string{"AiEnabledReqM", string(data)}, " ")
}
