package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTenantUserConfigurationResponse Response Object
type SetTenantUserConfigurationResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetTenantUserConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTenantUserConfigurationResponse struct{}"
	}

	return strings.Join([]string{"SetTenantUserConfigurationResponse", string(data)}, " ")
}
