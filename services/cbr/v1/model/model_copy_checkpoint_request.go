package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CopyCheckpointRequest struct {
	Body *CheckpointReplicateReq `json:"body,omitempty"`
}

func (o CopyCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyCheckpointRequest struct{}"
	}

	return strings.Join([]string{"CopyCheckpointRequest", string(data)}, " ")
}
