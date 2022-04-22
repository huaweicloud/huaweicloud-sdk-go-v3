package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EncryptDatakeyRequest struct {

	// API版本号
	VersionId string `json:"version_id"`

	Body *EncryptDatakeyRequestBody `json:"body,omitempty"`
}

func (o EncryptDatakeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptDatakeyRequest struct{}"
	}

	return strings.Join([]string{"EncryptDatakeyRequest", string(data)}, " ")
}
