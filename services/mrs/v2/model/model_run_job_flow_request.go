package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunJobFlowRequest struct {
	Body *RunJobFlowCommand `json:"body,omitempty"`
}

func (o RunJobFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunJobFlowRequest struct{}"
	}

	return strings.Join([]string{"RunJobFlowRequest", string(data)}, " ")
}
