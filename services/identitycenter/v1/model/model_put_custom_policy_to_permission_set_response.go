package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutCustomPolicyToPermissionSetResponse Response Object
type PutCustomPolicyToPermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PutCustomPolicyToPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutCustomPolicyToPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"PutCustomPolicyToPermissionSetResponse", string(data)}, " ")
}
