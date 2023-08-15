package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachManagedPolicyFromPermissionSetResponse Response Object
type DetachManagedPolicyFromPermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachManagedPolicyFromPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachManagedPolicyFromPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"DetachManagedPolicyFromPermissionSetResponse", string(data)}, " ")
}
