package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceMetricRequest Request Object
type ShowInstanceMetricRequest struct {
	Body *ShowInstanceMetricRequestBody `json:"body,omitempty"`
}

func (o ShowInstanceMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMetricRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceMetricRequest", string(data)}, " ")
}
