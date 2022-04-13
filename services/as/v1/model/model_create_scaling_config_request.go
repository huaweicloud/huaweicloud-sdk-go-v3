package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateScalingConfigRequest struct {
	Body *CreateScalingConfigOption `json:"body,omitempty"`
}

func (o CreateScalingConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingConfigRequest", string(data)}, " ")
}
