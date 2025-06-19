package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNewJobRequest Request Object
type UpdateNewJobRequest struct {
	Body *UpdateBuildJobRequestBody `json:"body,omitempty"`
}

func (o UpdateNewJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNewJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateNewJobRequest", string(data)}, " ")
}
