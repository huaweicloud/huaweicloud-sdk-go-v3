package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommitSimpleResultDto struct {

	// 提交列表
	Commits *[]CommitDto `json:"commits,omitempty"`
}

func (o CommitSimpleResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitSimpleResultDto struct{}"
	}

	return strings.Join([]string{"CommitSimpleResultDto", string(data)}, " ")
}
