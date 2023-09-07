package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseConnectionInfo 连接配置内容
type BaseConnectionInfo struct {
	AuthConfig *interface{} `json:"auth_config,omitempty"`

	AuthConfigId *string `json:"auth_config_id,omitempty"`

	AuthDynamic *interface{} `json:"auth_dynamic,omitempty"`

	AuthInfo *interface{} `json:"auth_info,omitempty"`

	AuthProp *interface{} `json:"auth_prop,omitempty"`

	AuthType *string `json:"auth_type,omitempty"`

	CdmParamsConfig *interface{} `json:"cdm_params_config,omitempty"`

	ConnectionName *string `json:"connection_name,omitempty"`

	ConnectorId *string `json:"connector_id,omitempty"`

	Description *string `json:"description,omitempty"`

	HostConfig *interface{} `json:"host_config,omitempty"`
}

func (o BaseConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseConnectionInfo struct{}"
	}

	return strings.Join([]string{"BaseConnectionInfo", string(data)}, " ")
}
