package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoBlockRes **参数解释**： 是否自动阻断告警 **取值范围**： - 0：不自动阻断告警 - 1：自动阻断告警
type AutoBlockRes struct {
}

func (o AutoBlockRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoBlockRes struct{}"
	}

	return strings.Join([]string{"AutoBlockRes", string(data)}, " ")
}
