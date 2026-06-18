package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepoLastStatisticsResponse Response Object
type ShowRepoLastStatisticsResponse struct {
	Event *StatisticEventsDto `json:"event,omitempty"`

	// **参数解释：** 统计信息数量 **取值范围：** 最小0 **默认取值：** 0
	Total *int32 `json:"total,omitempty"`

	// 统计信息
	Statistics *[]StatisticDto `json:"statistics,omitempty"`

	// 仓库近15日每日代码提交增减行数信息。
	Codelines *[]CodelineDto `json:"codelines,omitempty"`

	// **参数解释：** 分支提交总数。 **取值范围：** 最小0 **默认取值：** 0
	Count *int32 `json:"count,omitempty"`

	// **参数解释：** 仓库提交总数。 **取值范围：** 最小0 **默认取值：** 0
	AllBranchCommitsCount *int32 `json:"all_branch_commits_count,omitempty"`
	HttpStatusCode        int    `json:"-"`
}

func (o ShowRepoLastStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoLastStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowRepoLastStatisticsResponse", string(data)}, " ")
}
