package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteGenerateReportRequest struct {
	Body *ExecuteGenerateReportRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ExecuteGenerateReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGenerateReportRequest struct{}"
	}

	return strings.Join([]string{"ExecuteGenerateReportRequest", string(data)}, " ")
}
