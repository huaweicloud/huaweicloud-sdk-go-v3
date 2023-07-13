package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiagnoseJobRequest Request Object
type ListDiagnoseJobRequest struct {

	// 任务id
	JobId *string `json:"job_id,omitempty"`
}

func (o ListDiagnoseJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnoseJobRequest struct{}"
	}

	return strings.Join([]string{"ListDiagnoseJobRequest", string(data)}, " ")
}
