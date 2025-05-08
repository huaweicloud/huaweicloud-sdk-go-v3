package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObjectReplicationRequest Request Object
type CreateObjectReplicationRequest struct {
	Body *CreateObjectReplicationRequestBody `json:"body,omitempty"`
}

func (o CreateObjectReplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObjectReplicationRequest struct{}"
	}

	return strings.Join([]string{"CreateObjectReplicationRequest", string(data)}, " ")
}
