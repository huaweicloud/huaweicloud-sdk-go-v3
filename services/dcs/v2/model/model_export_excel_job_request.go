package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportExcelJobRequest Request Object
type ExportExcelJobRequest struct {

	// jobID
	JobId string `json:"job_id"`
}

func (o ExportExcelJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportExcelJobRequest struct{}"
	}

	return strings.Join([]string{"ExportExcelJobRequest", string(data)}, " ")
}
