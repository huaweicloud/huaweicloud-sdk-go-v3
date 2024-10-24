package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomPolicyFromPermissionSetResponse Response Object
type DeleteCustomPolicyFromPermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomPolicyFromPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomPolicyFromPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomPolicyFromPermissionSetResponse", string(data)}, " ")
}
