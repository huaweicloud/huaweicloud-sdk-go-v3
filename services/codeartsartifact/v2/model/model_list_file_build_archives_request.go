package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFileBuildArchivesRequest Request Object
type ListFileBuildArchivesRequest struct {

	// 父节点id
	ParentId *string `json:"parent_id,omitempty"`

	// 构建id
	BuildId *string `json:"build_id,omitempty"`

	// 构建序号
	BuildNo *string `json:"build_no,omitempty"`

	// 页数
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页数据量
	PageSize *int32 `json:"page_size,omitempty"`

	// codehub仓库分支
	RepoBranch *string `json:"repo_branch,omitempty"`
}

func (o ListFileBuildArchivesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFileBuildArchivesRequest struct{}"
	}

	return strings.Join([]string{"ListFileBuildArchivesRequest", string(data)}, " ")
}
