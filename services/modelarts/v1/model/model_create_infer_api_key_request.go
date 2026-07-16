package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInferApiKeyRequest Request Object
type CreateInferApiKeyRequest struct {
	Body *CreateApiKeyReq `json:"body,omitempty"`
}

func (o CreateInferApiKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInferApiKeyRequest struct{}"
	}

	return strings.Join([]string{"CreateInferApiKeyRequest", string(data)}, " ")
}
