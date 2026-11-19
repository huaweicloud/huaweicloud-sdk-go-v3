package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotaHostStatisticsResponse Response Object
type ShowQuotaHostStatisticsResponse struct {

	// **参数解释**: 主机总数量 **取值范围**: 最小值0，最大值2147483647
	HostTotalNum *int32 `json:"host_total_num,omitempty"`

	// **参数解释**: 配额信息列表，包含各版本的配额统计 **取值范围**: 0-200个QuotaHostStatistics对象
	DataList       *[]QuotaHostStatistics `json:"data_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowQuotaHostStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaHostStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaHostStatisticsResponse", string(data)}, " ")
}
