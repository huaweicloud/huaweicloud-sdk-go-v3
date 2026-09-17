package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExceptionMetricsRequest Request Object
type ListExceptionMetricsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ListExceptionMetricsRequestBody `json:"body,omitempty"`
}

func (o ListExceptionMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExceptionMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListExceptionMetricsRequest", string(data)}, " ")
}
