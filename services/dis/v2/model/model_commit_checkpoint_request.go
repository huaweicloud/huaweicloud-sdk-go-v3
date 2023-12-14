package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitCheckpointRequest Request Object
type CommitCheckpointRequest struct {
	Body *CommitCheckpointReq `json:"body,omitempty"`
}

func (o CommitCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitCheckpointRequest struct{}"
	}

	return strings.Join([]string{"CommitCheckpointRequest", string(data)}, " ")
}
