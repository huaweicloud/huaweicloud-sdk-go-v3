package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateManualBackupRequest struct {
	Body *CreateManualBackupRequestBody `json:"body,omitempty"`
}

func (o CreateManualBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualBackupRequest struct{}"
	}

	return strings.Join([]string{"CreateManualBackupRequest", string(data)}, " ")
}
