package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInferTempApiKeyRequest Request Object
type CreateInferTempApiKeyRequest struct {
	Body *CreateTempApiKeyReq `json:"body,omitempty"`
}

func (o CreateInferTempApiKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInferTempApiKeyRequest struct{}"
	}

	return strings.Join([]string{"CreateInferTempApiKeyRequest", string(data)}, " ")
}
