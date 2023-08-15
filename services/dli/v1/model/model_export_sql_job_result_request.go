package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSqlJobResultRequest Request Object
type ExportSqlJobResultRequest struct {

	// 作业ID
	JobId string `json:"job_id"`

	Body *ExportSqlResultRequestBody `json:"body,omitempty"`
}

func (o ExportSqlJobResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSqlJobResultRequest struct{}"
	}

	return strings.Join([]string{"ExportSqlJobResultRequest", string(data)}, " ")
}
