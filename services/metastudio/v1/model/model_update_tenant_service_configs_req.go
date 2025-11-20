package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTenantServiceConfigsReq 租户业务配置
type UpdateTenantServiceConfigsReq struct {
	SubAccountControlConfig *SubAccountControlConfig `json:"sub_account_control_config,omitempty"`

	// AI标识开关
	IsAiMarkOn *bool `json:"is_ai_mark_on,omitempty"`
}

func (o UpdateTenantServiceConfigsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantServiceConfigsReq struct{}"
	}

	return strings.Join([]string{"UpdateTenantServiceConfigsReq", string(data)}, " ")
}
