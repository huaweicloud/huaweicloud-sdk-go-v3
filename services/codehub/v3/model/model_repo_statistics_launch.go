package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoStatisticsLaunch struct {

	// 仓库是否可以统计
	CanStatistics *bool `json:"can_statistics,omitempty" xml:"can_statistics"`

	// sidekiq任务的 id
	JoinId *string `json:"join_id,omitempty" xml:"join_id"`

	// 启动仓库统计返回的信息
	Message *string `json:"message,omitempty" xml:"message"`
}

func (o RepoStatisticsLaunch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoStatisticsLaunch struct{}"
	}

	return strings.Join([]string{"RepoStatisticsLaunch", string(data)}, " ")
}
