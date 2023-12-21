package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterAuthorizationRequest Request Object
type RegisterAuthorizationRequest struct {
	Body *AuthorizeCsmsAndKmsRequestBody `json:"body,omitempty"`
}

func (o RegisterAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"RegisterAuthorizationRequest", string(data)}, " ")
}
