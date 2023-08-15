package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAndExecuteJobRequest Request Object
type CreateAndExecuteJobRequest struct {
	Body *SubmitJobReqV11 `json:"body,omitempty"`
}

func (o CreateAndExecuteJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAndExecuteJobRequest struct{}"
	}

	return strings.Join([]string{"CreateAndExecuteJobRequest", string(data)}, " ")
}
