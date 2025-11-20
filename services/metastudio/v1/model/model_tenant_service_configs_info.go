package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantServiceConfigsInfo 租户业务配置
type TenantServiceConfigsInfo struct {
	ServiceSharedConfig *ServiceSharedConfig `json:"service_shared_config,omitempty"`

	TenantLogConfig *TenantLogConfig `json:"tenant_log_config,omitempty"`

	SubAccountControlConfig *SubAccountControlConfig `json:"sub_account_control_config,omitempty"`

	// AI标识开关
	IsAiMarkOn *bool `json:"is_ai_mark_on,omitempty"`
}

func (o TenantServiceConfigsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantServiceConfigsInfo struct{}"
	}

	return strings.Join([]string{"TenantServiceConfigsInfo", string(data)}, " ")
}
