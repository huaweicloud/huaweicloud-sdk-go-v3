package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNewJobRequest Request Object
type CreateNewJobRequest struct {
	Body *CreateBuildJobRequestBody `json:"body,omitempty"`
}

func (o CreateNewJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNewJobRequest struct{}"
	}

	return strings.Join([]string{"CreateNewJobRequest", string(data)}, " ")
}
