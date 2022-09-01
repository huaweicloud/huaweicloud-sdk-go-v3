package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowImageBlobRequest struct {

	// 仓库id
	RepositoryUuid string `json:"repository_uuid" xml:"repository_uuid"`

	// 分支名称
	BranchName string `json:"branch_name" xml:"branch_name"`

	// 图片路径
	Path string `json:"path" xml:"path"`
}

func (o ShowImageBlobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageBlobRequest struct{}"
	}

	return strings.Join([]string{"ShowImageBlobRequest", string(data)}, " ")
}
