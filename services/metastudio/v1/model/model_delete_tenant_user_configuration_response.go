package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTenantUserConfigurationResponse Response Object
type DeleteTenantUserConfigurationResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTenantUserConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTenantUserConfigurationResponse struct{}"
	}

	return strings.Join([]string{"DeleteTenantUserConfigurationResponse", string(data)}, " ")
}
