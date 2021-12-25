package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 变更行数
type SpecificCommitInfoStats struct {
	// 变更增加的行数

	Additions *int32 `json:"additions,omitempty"`
	// 变更删除的行数

	Deletions *int32 `json:"deletions,omitempty"`
	// 变更的总行数

	Total *int32 `json:"total,omitempty"`
}

func (o SpecificCommitInfoStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecificCommitInfoStats struct{}"
	}

	return strings.Join([]string{"SpecificCommitInfoStats", string(data)}, " ")
}
