package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommitStatistic struct {

	// 增加的行数
	Additions *int32 `json:"additions,omitempty"`

	// 删除的行数
	Deletions *int32 `json:"deletions,omitempty"`
}

func (o CommitStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitStatistic struct{}"
	}

	return strings.Join([]string{"CommitStatistic", string(data)}, " ")
}
