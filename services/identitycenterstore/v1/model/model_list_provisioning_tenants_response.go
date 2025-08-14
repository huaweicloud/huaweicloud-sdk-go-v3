package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProvisioningTenantsResponse Response Object
type ListProvisioningTenantsResponse struct {

	// SCIM自动预配配置信息
	ProvisioningTenants *[]ProvisioningTenant `json:"provisioning_tenants,omitempty"`
	HttpStatusCode      int                   `json:"-"`
}

func (o ListProvisioningTenantsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProvisioningTenantsResponse struct{}"
	}

	return strings.Join([]string{"ListProvisioningTenantsResponse", string(data)}, " ")
}
