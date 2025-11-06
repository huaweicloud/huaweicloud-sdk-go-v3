package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommitV2 struct {

	// 提交时间
	CommittedDate *string `json:"committed_date,omitempty"`

	// 提交者
	CommitterName *string `json:"committer_name,omitempty"`

	// 提交id
	Id *string `json:"id,omitempty"`

	// 提交信息
	Message *string `json:"message,omitempty"`

	// 提交短id
	ShortId *string `json:"short_id,omitempty"`

	// 提交标题
	Title *string `json:"title,omitempty"`
}

func (o CommitV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitV2 struct{}"
	}

	return strings.Join([]string{"CommitV2", string(data)}, " ")
}
