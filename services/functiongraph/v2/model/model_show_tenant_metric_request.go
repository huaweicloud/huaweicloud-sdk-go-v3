package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTenantMetricRequest struct {

	// 时间段，单位为分钟
	Period *string `json:"period,omitempty" xml:"period"`

	// 开始时间，精确到ms的时间戳
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 结束时间，精确到ms的时间戳
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`
}

func (o ShowTenantMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantMetricRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantMetricRequest", string(data)}, " ")
}
