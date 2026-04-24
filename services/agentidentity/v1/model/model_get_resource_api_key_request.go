package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetResourceApiKeyRequest Request Object
type GetResourceApiKeyRequest struct {
	Body *GetResourceApiKeyRequestBody `json:"body,omitempty"`
}

func (o GetResourceApiKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceApiKeyRequest struct{}"
	}

	return strings.Join([]string{"GetResourceApiKeyRequest", string(data)}, " ")
}
