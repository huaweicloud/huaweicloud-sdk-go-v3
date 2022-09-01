package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCommitsByBranchRequest struct {

	// 仓库组名
	GroupName string `json:"group_name" xml:"group_name"`

	// 分页索引
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index"`

	// 分页索引
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size"`

	// 分支或标签名，支持SHA格式
	RefName string `json:"ref_name" xml:"ref_name"`

	// 仓库名
	RepositoryName string `json:"repository_name" xml:"repository_name"`
}

func (o ShowCommitsByBranchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommitsByBranchRequest struct{}"
	}

	return strings.Join([]string{"ShowCommitsByBranchRequest", string(data)}, " ")
}
