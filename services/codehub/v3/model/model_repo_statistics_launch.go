package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoStatisticsLaunch struct {
	// 仓库是否可以统计

	CanStatistics *bool `json:"can_statistics,omitempty"`
	// sidekiq任务的 id

	JoinId *string `json:"join_id,omitempty"`
	// 启动仓库统计返回的信息

	Message *string `json:"message,omitempty"`
}

func (o RepoStatisticsLaunch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoStatisticsLaunch struct{}"
	}

	return strings.Join([]string{"RepoStatisticsLaunch", string(data)}, " ")
}
