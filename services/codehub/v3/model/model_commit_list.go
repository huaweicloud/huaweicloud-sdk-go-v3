package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommitList struct {

	// 提交列表
	Commits *[]Commit `json:"commits,omitempty"`

	// 提交总数
	Total *int32 `json:"total,omitempty"`
}

func (o CommitList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitList struct{}"
	}

	return strings.Join([]string{"CommitList", string(data)}, " ")
}
