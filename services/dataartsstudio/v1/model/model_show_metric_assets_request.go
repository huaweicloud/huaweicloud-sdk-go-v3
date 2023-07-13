package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetricAssetsRequest Request Object
type ShowMetricAssetsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *MetricOpenSearchParams `json:"body,omitempty"`
}

func (o ShowMetricAssetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricAssetsRequest struct{}"
	}

	return strings.Join([]string{"ShowMetricAssetsRequest", string(data)}, " ")
}
