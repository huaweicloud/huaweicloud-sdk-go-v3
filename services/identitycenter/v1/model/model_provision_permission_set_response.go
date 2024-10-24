package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProvisionPermissionSetResponse Response Object
type ProvisionPermissionSetResponse struct {
	PermissionSetProvisioningStatus *PermissionSetProvisioningStatusDto `json:"permission_set_provisioning_status,omitempty"`
	HttpStatusCode                  int                                 `json:"-"`
}

func (o ProvisionPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProvisionPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"ProvisionPermissionSetResponse", string(data)}, " ")
}
