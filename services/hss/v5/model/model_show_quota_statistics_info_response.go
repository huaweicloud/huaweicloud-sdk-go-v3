package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotaStatisticsInfoResponse Response Object
type ShowQuotaStatisticsInfoResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 防护配额统计列表
	DataList       *[]QuotaStatisticsInfo `json:"data_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowQuotaStatisticsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaStatisticsInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaStatisticsInfoResponse", string(data)}, " ")
}
