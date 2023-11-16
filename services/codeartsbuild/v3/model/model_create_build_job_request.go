package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBuildJobRequest Request Object
type CreateBuildJobRequest struct {
	Body *CreateBuildJobRequestBody `json:"body,omitempty"`
}

func (o CreateBuildJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBuildJobRequest struct{}"
	}

	return strings.Join([]string{"CreateBuildJobRequest", string(data)}, " ")
}
