package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAuthorizationRequest Request Object
type ChangeAuthorizationRequest struct {
	Body *GrantDataPermissionReq `json:"body,omitempty"`
}

func (o ChangeAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"ChangeAuthorizationRequest", string(data)}, " ")
}
