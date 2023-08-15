package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteGenerateReportRequest Request Object
type ExecuteGenerateReportRequest struct {
	Body *ExecuteGenerateReportRequestBody `json:"body,omitempty"`
}

func (o ExecuteGenerateReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGenerateReportRequest struct{}"
	}

	return strings.Join([]string{"ExecuteGenerateReportRequest", string(data)}, " ")
}
