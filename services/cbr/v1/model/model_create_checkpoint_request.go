package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCheckpointRequest Request Object
type CreateCheckpointRequest struct {
	Body *VaultBackupReq `json:"body,omitempty"`
}

func (o CreateCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCheckpointRequest struct{}"
	}

	return strings.Join([]string{"CreateCheckpointRequest", string(data)}, " ")
}
