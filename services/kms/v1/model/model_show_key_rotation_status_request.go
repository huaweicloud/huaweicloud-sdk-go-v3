package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowKeyRotationStatusRequest struct {

	// API版本号
	VersionId string `json:"version_id"`

	Body *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o ShowKeyRotationStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyRotationStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowKeyRotationStatusRequest", string(data)}, " ")
}
