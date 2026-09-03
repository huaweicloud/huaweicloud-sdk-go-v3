package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceMetricResponse Response Object
type ShowInstanceMetricResponse struct {

	// 实例指标信息列表
	DasMetricInfos *[]DasMetricInfo `json:"das_metric_infos,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowInstanceMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMetricResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceMetricResponse", string(data)}, " ")
}
