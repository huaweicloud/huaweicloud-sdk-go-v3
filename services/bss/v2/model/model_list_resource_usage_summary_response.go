package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceUsageSummaryResponse Response Object
type ListResourceUsageSummaryResponse struct {

	// 总条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 统计值，按照资源ID维度返回的月度统计结果。具体请参见表3。  说明： 每月2日20点后可查询上月数据；若查询当月数据，则只返回资源ID。
	SummaryUsageInfoList *[]StatUsageSummaryInfo `json:"summary_usage_info_list,omitempty"`
	HttpStatusCode       int                     `json:"-"`
}

func (o ListResourceUsageSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceUsageSummaryResponse struct{}"
	}

	return strings.Join([]string{"ListResourceUsageSummaryResponse", string(data)}, " ")
}
