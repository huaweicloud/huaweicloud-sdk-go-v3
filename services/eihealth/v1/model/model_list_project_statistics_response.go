package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectStatisticsResponse Response Object
type ListProjectStatisticsResponse struct {

	// 空间统计信息。
	Projects *[]ProjectStatistic `json:"projects,omitempty"`

	// 空间总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectStatisticsResponse", string(data)}, " ")
}
