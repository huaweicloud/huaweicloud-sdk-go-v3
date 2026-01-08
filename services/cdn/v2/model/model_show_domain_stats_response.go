package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainStatsResponse Response Object
type ShowDomainStatsResponse struct {

	// 查询起始时间戳。
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询结束时间戳
	EndTime *int64 `json:"end_time,omitempty"`

	// 参数类型支持：flux(流量)，req_num(请求总数)。
	StatType *string `json:"stat_type,omitempty"`

	// **参数解释：** 规则行为 **约束限制：** 不涉及
	Action *string `json:"action,omitempty"`

	// 查询时间间隔，单位：秒
	Interval *int64 `json:"interval,omitempty"`

	// 按指定的分组方式组织的数据
	Result         map[string]interface{} `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowDomainStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainStatsResponse", string(data)}, " ")
}
