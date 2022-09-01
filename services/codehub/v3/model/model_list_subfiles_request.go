package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSubfilesRequest struct {

	// 仓库id
	RepositoryUuid string `json:"repository_uuid" xml:"repository_uuid"`

	// 分支名称
	BranchName string `json:"branch_name" xml:"branch_name"`

	// 文件路径
	Path *string `json:"path,omitempty" xml:"path"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 记录数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListSubfilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubfilesRequest struct{}"
	}

	return strings.Join([]string{"ListSubfilesRequest", string(data)}, " ")
}
