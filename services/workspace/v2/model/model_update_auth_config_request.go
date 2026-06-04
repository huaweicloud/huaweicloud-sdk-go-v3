package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthConfigRequest Request Object
type UpdateAuthConfigRequest struct {

	// 认证配置ID
	AuthConfigId string `json:"auth_config_id"`

	Body *UpdateAuthConfigReq `json:"body,omitempty"`
}

func (o UpdateAuthConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateAuthConfigRequest", string(data)}, " ")
}
