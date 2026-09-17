package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSingleMetricRequest Request Object
type ShowSingleMetricRequest struct {
	Body *ShowSingleMetricRequestBody `json:"body,omitempty"`
}

func (o ShowSingleMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleMetricRequest struct{}"
	}

	return strings.Join([]string{"ShowSingleMetricRequest", string(data)}, " ")
}
