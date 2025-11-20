package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantServiceConfigsResponse Response Object
type ShowTenantServiceConfigsResponse struct {
	ServiceSharedConfig *ServiceSharedConfig `json:"service_shared_config,omitempty"`

	TenantLogConfig *TenantLogConfig `json:"tenant_log_config,omitempty"`

	SubAccountControlConfig *SubAccountControlConfig `json:"sub_account_control_config,omitempty"`

	// AI标识开关
	IsAiMarkOn *bool `json:"is_ai_mark_on,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTenantServiceConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantServiceConfigsResponse struct{}"
	}

	return strings.Join([]string{"ShowTenantServiceConfigsResponse", string(data)}, " ")
}
