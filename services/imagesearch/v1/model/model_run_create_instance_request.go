package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunCreateInstanceRequest struct {
	Body *CreateInstanceReq `json:"body,omitempty"`
}

func (o RunCreateInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCreateInstanceRequest struct{}"
	}

	return strings.Join([]string{"RunCreateInstanceRequest", string(data)}, " ")
}
