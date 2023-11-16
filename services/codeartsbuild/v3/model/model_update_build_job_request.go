package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBuildJobRequest Request Object
type UpdateBuildJobRequest struct {
	Body *UpdateBuildJobRequestBody `json:"body,omitempty"`
}

func (o UpdateBuildJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBuildJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateBuildJobRequest", string(data)}, " ")
}
