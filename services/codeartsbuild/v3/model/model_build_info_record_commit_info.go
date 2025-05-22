package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BuildInfoRecordCommitInfo 代码提交信息
type BuildInfoRecordCommitInfo struct {

	// 代码提交的commit id
	CommitId *string `json:"commit_id,omitempty"`

	// 提交时间
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o BuildInfoRecordCommitInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildInfoRecordCommitInfo struct{}"
	}

	return strings.Join([]string{"BuildInfoRecordCommitInfo", string(data)}, " ")
}
