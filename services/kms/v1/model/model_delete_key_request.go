package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteKeyRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *ScheduleKeyDeletionRequestBody `json:"body,omitempty"`
}

func (o DeleteKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeyRequest struct{}"
	}

	return strings.Join([]string{"DeleteKeyRequest", string(data)}, " ")
}
