package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EncryptDataRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *EncryptDataRequestBody `json:"body,omitempty"`
}

func (o EncryptDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDataRequest struct{}"
	}

	return strings.Join([]string{"EncryptDataRequest", string(data)}, " ")
}
