package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CommitCheckpointRequest struct {
	Body *CommitCheckpointRequestBody `json:"body,omitempty"`
}

func (o CommitCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitCheckpointRequest struct{}"
	}

	return strings.Join([]string{"CommitCheckpointRequest", string(data)}, " ")
}
