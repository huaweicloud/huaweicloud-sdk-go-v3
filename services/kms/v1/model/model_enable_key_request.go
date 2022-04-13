package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EnableKeyRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o EnableKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyRequest struct{}"
	}

	return strings.Join([]string{"EnableKeyRequest", string(data)}, " ")
}
