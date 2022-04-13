package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMetricValueRequest struct {
	// 资产ID

	AssetId string `json:"asset_id"`

	Body *GetMetricsValue `json:"body,omitempty"`
}

func (o ShowMetricValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricValueRequest struct{}"
	}

	return strings.Join([]string{"ShowMetricValueRequest", string(data)}, " ")
}
