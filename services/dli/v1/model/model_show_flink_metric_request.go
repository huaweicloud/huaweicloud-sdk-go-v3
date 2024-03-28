package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkMetricRequest Request Object
type ShowFlinkMetricRequest struct {
	Body *ShowFlinkMetricRequestBody `json:"body,omitempty"`
}

func (o ShowFlinkMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkMetricRequest struct{}"
	}

	return strings.Join([]string{"ShowFlinkMetricRequest", string(data)}, " ")
}
