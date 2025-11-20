package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantServiceConfigsResponse Response Object
type UpdateTenantServiceConfigsResponse struct {
	ServiceSharedConfig *ServiceSharedConfig `json:"service_shared_config,omitempty"`

	TenantLogConfig *TenantLogConfig `json:"tenant_log_config,omitempty"`

	SubAccountControlConfig *SubAccountControlConfig `json:"sub_account_control_config,omitempty"`

	// AI标识开关
	IsAiMarkOn *bool `json:"is_ai_mark_on,omitempty"`

	// 租户project ID
	TenantId *string `json:"tenant_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTenantServiceConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantServiceConfigsResponse struct{}"
	}

	return strings.Join([]string{"UpdateTenantServiceConfigsResponse", string(data)}, " ")
}
