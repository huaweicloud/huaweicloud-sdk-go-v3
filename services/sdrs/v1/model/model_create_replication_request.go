package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateReplicationRequest struct {
	Body *CreateReplicationRequestBody `json:"body,omitempty"`
}

func (o CreateReplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReplicationRequest struct{}"
	}

	return strings.Join([]string{"CreateReplicationRequest", string(data)}, " ")
}
