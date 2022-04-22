package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStatisticalDataResponse struct {

	// 仓库名称
	RepoName *string `json:"repoName,omitempty"`

	// 提交次数
	CommitCount *int32 `json:"commitCount,omitempty"`

	// 仓库容量
	RepoSize *string `json:"repoSize,omitempty"`

	// 最近一次提交时间
	LastCommitTime *string `json:"lastCommitTime,omitempty"`

	// 代码行数
	CodeLines *int32 `json:"codeLines,omitempty"`

	// 分支数量
	BranchCount *int32 `json:"branchCount,omitempty"`

	// 代码仓下载地址
	ArchiveUrl     *string `json:"archiveUrl,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStatisticalDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticalDataResponse struct{}"
	}

	return strings.Join([]string{"ShowStatisticalDataResponse", string(data)}, " ")
}
