package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowRepositoryStatisticsRequestBody struct {

	// 仓库分支名
	BranchName string `json:"branch_name"`
}

func (o ShowRepositoryStatisticsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryStatisticsRequestBody struct{}"
	}

	return strings.Join([]string{"ShowRepositoryStatisticsRequestBody", string(data)}, " ")
}
