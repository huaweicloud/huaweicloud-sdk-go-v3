package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTenantMetricRequest struct {
	// 时间段，单位为分钟

	Period *string `json:"period,omitempty"`
}

func (o ShowTenantMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantMetricRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantMetricRequest", string(data)}, " ")
}
