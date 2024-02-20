package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ArchiveRuleId 存档规则的唯一标识符。
type ArchiveRuleId struct {
}

func (o ArchiveRuleId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArchiveRuleId struct{}"
	}

	return strings.Join([]string{"ArchiveRuleId", string(data)}, " ")
}
