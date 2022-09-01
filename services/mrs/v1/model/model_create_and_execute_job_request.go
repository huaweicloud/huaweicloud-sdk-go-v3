package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAndExecuteJobRequest struct {
	Body *SubmitJobReqV11 `json:"body,omitempty" xml:"body"`
}

func (o CreateAndExecuteJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAndExecuteJobRequest struct{}"
	}

	return strings.Join([]string{"CreateAndExecuteJobRequest", string(data)}, " ")
}
