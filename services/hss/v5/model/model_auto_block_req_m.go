package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoBlockReqM **参数解释**： 是否自动阻断告警 **约束限制**： 必填 **取值范围**： - 0：不自动阻断告警 - 1：自动阻断告警  **默认取值**： 不涉及
type AutoBlockReqM struct {
}

func (o AutoBlockReqM) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoBlockReqM struct{}"
	}

	return strings.Join([]string{"AutoBlockReqM", string(data)}, " ")
}
