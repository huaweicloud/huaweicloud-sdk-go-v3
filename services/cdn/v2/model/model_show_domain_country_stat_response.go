package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainCountryStatResponse Response Object
type ShowDomainCountryStatResponse struct {

	// **参数解释：** 规则行为 **约束限制：** 不涉及
	Action *string `json:"action,omitempty"`

	// 查询起始时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询结束时间，相对于UTC 1970-01-01到当前时间相隔的毫秒数。
	EndTime *int64 `json:"end_time,omitempty"`

	// 参数类型支持：flux(流量)，req_num(请求总数)。
	StatType       *string `json:"stat_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDomainCountryStatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainCountryStatResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainCountryStatResponse", string(data)}, " ")
}
