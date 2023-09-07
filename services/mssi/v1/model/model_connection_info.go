package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionInfo 连接配置内容
type ConnectionInfo struct {
	AuthConfig *interface{} `json:"auth_config,omitempty"`

	AuthConfigId *string `json:"auth_config_id,omitempty"`

	AuthDynamic *interface{} `json:"auth_dynamic,omitempty"`

	AuthId *string `json:"auth_id,omitempty"`

	AuthInfo *interface{} `json:"auth_info,omitempty"`

	AuthKey *string `json:"auth_key,omitempty"`

	AuthName *string `json:"auth_name,omitempty"`

	AuthProp *interface{} `json:"auth_prop,omitempty"`

	AuthType *string `json:"auth_type,omitempty"`

	CdmParamsConfig *interface{} `json:"cdm_params_config,omitempty"`

	ConnectionName *string `json:"connection_name,omitempty"`

	ConnectorId *string `json:"connector_id,omitempty"`

	ConnectorName *string `json:"connector_name,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	CreatedBy *string `json:"created_by,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainId *string `json:"domain_id,omitempty"`

	HostConfig *interface{} `json:"host_config,omitempty"`

	Id *string `json:"id,omitempty"`

	IsOpen *int32 `json:"is_open,omitempty"`

	Logo *string `json:"logo,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Status *string `json:"status,omitempty"`

	Type *string `json:"type,omitempty"`

	UpdatedBy *string `json:"updated_by,omitempty"`

	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`

	UserId *string `json:"user_id,omitempty"`
}

func (o ConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionInfo struct{}"
	}

	return strings.Join([]string{"ConnectionInfo", string(data)}, " ")
}
