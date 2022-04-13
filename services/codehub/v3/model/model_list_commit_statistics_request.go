package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCommitStatisticsRequest struct {
	// 分支名

	BranchName string `json:"branch_name"`
	// 仓库的主键id

	RepositoryId string `json:"repository_id"`
}

func (o ListCommitStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListCommitStatisticsRequest", string(data)}, " ")
}
