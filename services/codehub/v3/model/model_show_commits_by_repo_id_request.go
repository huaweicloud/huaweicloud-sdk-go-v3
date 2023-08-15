package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommitsByRepoIdRequest Request Object
type ShowCommitsByRepoIdRequest struct {

	// 提交作者
	Author *string `json:"author,omitempty"`

	// 起始提交日期，格式为yyyy-MM-dd
	BeginDate *string `json:"begin_date,omitempty"`

	// 终止提交日期，格式为yyyy-MM-dd
	EndDate *string `json:"end_date,omitempty"`

	// 提交信息
	Message *string `json:"message,omitempty"`

	// 分页索引
	PageIndex *int32 `json:"page_index,omitempty"`

	// 每页数据量
	PageSize *int32 `json:"page_size,omitempty"`

	// 文件路径
	Path *string `json:"path,omitempty"`

	// 分支或标签名，支持SHA格式
	RefName string `json:"ref_name"`

	// 仓库主键id
	RepositoryId int32 `json:"repository_id"`

	// 提交的文件变更详情信息（不包含diff）
	StatFormat *string `json:"stat_format,omitempty"`
}

func (o ShowCommitsByRepoIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommitsByRepoIdRequest struct{}"
	}

	return strings.Join([]string{"ShowCommitsByRepoIdRequest", string(data)}, " ")
}
