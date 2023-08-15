package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribePermissionSetProvisioningStatusResponse Response Object
type DescribePermissionSetProvisioningStatusResponse struct {
	PermissionSetProvisioningStatus *PermissionSetProvisioningStatusDto `json:"permission_set_provisioning_status,omitempty"`
	HttpStatusCode                  int                                 `json:"-"`
}

func (o DescribePermissionSetProvisioningStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribePermissionSetProvisioningStatusResponse struct{}"
	}

	return strings.Join([]string{"DescribePermissionSetProvisioningStatusResponse", string(data)}, " ")
}
