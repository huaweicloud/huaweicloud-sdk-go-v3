package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMetricValueRequest struct {

	// 资产ID
	AssetId string `json:"asset_id" xml:"asset_id"`

	Body *GetMetricsValue `json:"body,omitempty" xml:"body"`
}

func (o ShowMetricValueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricValueRequest struct{}"
	}

	return strings.Join([]string{"ShowMetricValueRequest", string(data)}, " ")
}
