package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncExternalUserRequest Request Object
type SyncExternalUserRequest struct {

	// 认证配置id。
	AuthConfigId string `json:"auth_config_id"`

	// 域控id。
	DomainId string `json:"domain_id"`

	// 待同步的用户归属的企业项目id。
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

func (o SyncExternalUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncExternalUserRequest struct{}"
	}

	return strings.Join([]string{"SyncExternalUserRequest", string(data)}, " ")
}
