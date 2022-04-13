package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
