package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ArchiveRuleName 创建存档规则的名称。
type ArchiveRuleName struct {
}

func (o ArchiveRuleName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArchiveRuleName struct{}"
	}

	return strings.Join([]string{"ArchiveRuleName", string(data)}, " ")
}
