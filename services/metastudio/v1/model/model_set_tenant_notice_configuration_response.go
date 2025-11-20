package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTenantNoticeConfigurationResponse Response Object
type SetTenantNoticeConfigurationResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetTenantNoticeConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTenantNoticeConfigurationResponse struct{}"
	}

	return strings.Join([]string{"SetTenantNoticeConfigurationResponse", string(data)}, " ")
}
