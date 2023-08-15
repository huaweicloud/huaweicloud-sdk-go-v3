package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunShellCommandRequest Request Object
type RunShellCommandRequest struct {
	Body *RunShellCommandRequestBody `json:"body,omitempty"`
}

func (o RunShellCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunShellCommandRequest struct{}"
	}

	return strings.Join([]string{"RunShellCommandRequest", string(data)}, " ")
}
