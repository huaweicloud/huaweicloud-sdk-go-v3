package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListMetricDataRequest Request Object
type BatchListMetricDataRequest struct {
	Body *BatchListMetricDataRequestBody `json:"body,omitempty"`
}

func (o BatchListMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMetricDataRequest struct{}"
	}

	return strings.Join([]string{"BatchListMetricDataRequest", string(data)}, " ")
}
