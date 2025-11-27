package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSubscribeReportRequest Request Object
type BatchSubscribeReportRequest struct {
	Body *BatchSubscribeReportRequestBody `json:"body,omitempty"`
}

func (o BatchSubscribeReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSubscribeReportRequest struct{}"
	}

	return strings.Join([]string{"BatchSubscribeReportRequest", string(data)}, " ")
}
