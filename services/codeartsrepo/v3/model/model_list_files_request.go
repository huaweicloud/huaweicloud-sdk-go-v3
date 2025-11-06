package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFilesRequest Request Object
type ListFilesRequest struct {

	// 仓库id
	RepositoryUuid string `json:"repository_uuid"`

	// 分支名称
	BranchName string `json:"branch_name"`

	// 文件路径
	Path string `json:"path"`
}

func (o ListFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFilesRequest struct{}"
	}

	return strings.Join([]string{"ListFilesRequest", string(data)}, " ")
}
