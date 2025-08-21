package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommitStatisticsResponse Response Object
type ShowCommitStatisticsResponse struct {

	// **参数解释：** 提交统计。
	Commits *[]CommitStatisticsResultCommitDto `json:"commits,omitempty"`

	// **参数解释：** 提交人员统计。
	Statistics *[]StatisticDto `json:"statistics,omitempty"`

	// **参数解释：** 总数。
	Total *int32 `json:"total,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCommitStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommitStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowCommitStatisticsResponse", string(data)}, " ")
}
