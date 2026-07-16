package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthorizationRequest Request Object
type CreateAuthorizationRequest struct {
	Body *AuthorizationRequest `json:"body,omitempty"`
}

func (o CreateAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"CreateAuthorizationRequest", string(data)}, " ")
}
