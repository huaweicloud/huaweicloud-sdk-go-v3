package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnhanceFullSqlStatisticsResponse Response Object
type ListEnhanceFullSqlStatisticsResponse struct {

	// **参数解释**: 全量SQL聚合统计列表总数。 **取值范围**: 不涉及。
	TotalCount *int64 `json:"total_count,omitempty"`

	// **参数解释**: 全量SQL聚合统计列表。
	Statistics     *[]FullSqlStatisticInfoResult `json:"statistics,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListEnhanceFullSqlStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnhanceFullSqlStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListEnhanceFullSqlStatisticsResponse", string(data)}, " ")
}
