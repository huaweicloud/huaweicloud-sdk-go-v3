package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExecuteJobResponse Response Object
type CreateExecuteJobResponse struct {
	JobSubmitResult *JobSubmitResult `json:"job_submit_result,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o CreateExecuteJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExecuteJobResponse struct{}"
	}

	return strings.Join([]string{"CreateExecuteJobResponse", string(data)}, " ")
}
