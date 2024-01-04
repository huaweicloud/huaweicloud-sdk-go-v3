package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobAuditResultRequest Request Object
type ShowJobAuditResultRequest struct {

	// 任务id。
	JobId string `json:"job_id"`
}

func (o ShowJobAuditResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobAuditResultRequest struct{}"
	}

	return strings.Join([]string{"ShowJobAuditResultRequest", string(data)}, " ")
}
