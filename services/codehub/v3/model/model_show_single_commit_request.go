package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSingleCommitRequest struct {
	// 仓库短id

	RepoId int32 `json:"repo_id"`
	// commit id，仓库的branch名或tag名

	Sha string `json:"sha"`
	// 包括提交统计信息。默认值为true

	Stats *bool `json:"stats,omitempty"`
}

func (o ShowSingleCommitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleCommitRequest struct{}"
	}

	return strings.Join([]string{"ShowSingleCommitRequest", string(data)}, " ")
}
