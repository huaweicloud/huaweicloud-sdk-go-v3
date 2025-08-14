package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProvisioningTenantResponse Response Object
type DeleteProvisioningTenantResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProvisioningTenantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProvisioningTenantResponse struct{}"
	}

	return strings.Join([]string{"DeleteProvisioningTenantResponse", string(data)}, " ")
}
