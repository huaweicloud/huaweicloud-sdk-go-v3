package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ArchiveRuleUrn 存档规则的唯一资源标识符。
type ArchiveRuleUrn struct {
}

func (o ArchiveRuleUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArchiveRuleUrn struct{}"
	}

	return strings.Join([]string{"ArchiveRuleUrn", string(data)}, " ")
}
