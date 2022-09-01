package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateKeyRequest struct {

	// API版本号
	VersionId string `json:"version_id" xml:"version_id"`

	Body *CreateKeyRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyRequest struct{}"
	}

	return strings.Join([]string{"CreateKeyRequest", string(data)}, " ")
}
