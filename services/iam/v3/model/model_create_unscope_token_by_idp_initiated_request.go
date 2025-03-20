package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUnscopeTokenByIdpInitiatedRequest Request Object
type CreateUnscopeTokenByIdpInitiatedRequest struct {

	// 身份提供商ID。
	XIdpId string `json:"X-Idp-Id"`

	Body *CreateUnscopeTokenByIdpInitiatedRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CreateUnscopeTokenByIdpInitiatedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUnscopeTokenByIdpInitiatedRequest struct{}"
	}

	return strings.Join([]string{"CreateUnscopeTokenByIdpInitiatedRequest", string(data)}, " ")
}
