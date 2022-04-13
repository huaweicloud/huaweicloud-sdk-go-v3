package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunJobRequest struct {
	Body *RunJobRequestBody `json:"body,omitempty"`
}

func (o RunJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunJobRequest struct{}"
	}

	return strings.Join([]string{"RunJobRequest", string(data)}, " ")
}
