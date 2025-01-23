package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetricNamesSupportResponse Response Object
type ShowMetricNamesSupportResponse struct {

	// 支持指标名称列表
	SupportMetricNames *[]SupportMetricNameListSupportMetricNames `json:"support_metric_names,omitempty"`
	HttpStatusCode     int                                        `json:"-"`
}

func (o ShowMetricNamesSupportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricNamesSupportResponse struct{}"
	}

	return strings.Join([]string{"ShowMetricNamesSupportResponse", string(data)}, " ")
}
