package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScanJobRequest Request Object
type DeleteScanJobRequest struct {

	// 任务ID
	JobId string `json:"job_id"`
}

func (o DeleteScanJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScanJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteScanJobRequest", string(data)}, " ")
}
