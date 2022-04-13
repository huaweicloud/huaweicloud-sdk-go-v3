package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCommitsRequest struct {
	// 仓库短id

	RepoId int32 `json:"repo_id"`
	// 仓库的branch名或tag名，如果为空则查询默认分支

	RefName *string `json:"ref_name,omitempty"`
	// 在此日期之后或当天提交，格式 YYYY-MM-DDTHH:MM:SSZ

	Since *string `json:"since,omitempty"`
	// 在此日期之前或当天提交，格式 YYYY-MM-DDTHH:MM:SSZ

	Until *string `json:"until,omitempty"`
	// 文件路径

	Path *string `json:"path,omitempty"`
	// 是否检索仓库中每个提交

	All *bool `json:"all,omitempty"`
	// 有关每个提交的统计信息是否添加到响应中

	WithStats *bool `json:"with_stats,omitempty"`
}

func (o ListCommitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitsRequest struct{}"
	}

	return strings.Join([]string{"ListCommitsRequest", string(data)}, " ")
}
