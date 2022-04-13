package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRpoStatisticsResponse struct {
	// 资源的RPO超标趋势记录列表。

	ResourceRpoStatistics *[]RpoStattisticsParams `json:"resource_rpo_statistics,omitempty"`
	// 列表中包含的资源的RPO超标趋势记录个数。

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRpoStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRpoStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListRpoStatisticsResponse", string(data)}, " ")
}
