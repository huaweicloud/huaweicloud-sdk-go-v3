package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListSpecifiedMetricDataRequest Request Object
type BatchListSpecifiedMetricDataRequest struct {
	Body *BatchListSpecifiedMetricDataRequestBody `json:"body,omitempty"`
}

func (o BatchListSpecifiedMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListSpecifiedMetricDataRequest struct{}"
	}

	return strings.Join([]string{"BatchListSpecifiedMetricDataRequest", string(data)}, " ")
}
