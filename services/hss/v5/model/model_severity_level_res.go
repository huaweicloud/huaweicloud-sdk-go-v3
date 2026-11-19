package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SeverityLevelRes **参数解释**: 危险程度 **取值范围**: - Low：低危 - Medium：中危 - High：高危 - Critical：危急
type SeverityLevelRes struct {
}

func (o SeverityLevelRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SeverityLevelRes struct{}"
	}

	return strings.Join([]string{"SeverityLevelRes", string(data)}, " ")
}
