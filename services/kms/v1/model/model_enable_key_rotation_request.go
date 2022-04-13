package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EnableKeyRotationRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o EnableKeyRotationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyRotationRequest struct{}"
	}

	return strings.Join([]string{"EnableKeyRotationRequest", string(data)}, " ")
}
