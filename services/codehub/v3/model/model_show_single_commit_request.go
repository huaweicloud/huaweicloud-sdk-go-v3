package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSingleCommitRequest struct {

	// 仓库短id
	RepoId int32 `json:"repo_id" xml:"repo_id"`

	// commit id，仓库的branch名或tag名
	Sha string `json:"sha" xml:"sha"`

	// 包括提交统计信息。默认值为true
	Stats *bool `json:"stats,omitempty" xml:"stats"`
}

func (o ShowSingleCommitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleCommitRequest struct{}"
	}

	return strings.Join([]string{"ShowSingleCommitRequest", string(data)}, " ")
}
