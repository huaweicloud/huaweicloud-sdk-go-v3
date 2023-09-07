package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthConfigA struct {
	AuthConfig *interface{} `json:"auth_config,omitempty"`

	AuthDynamic *interface{} `json:"auth_dynamic,omitempty"`

	AuthInfo map[string]string `json:"auth_info,omitempty"`

	AuthProp *interface{} `json:"auth_prop,omitempty"`

	AuthType *string `json:"auth_type,omitempty"`

	CdmParamsConfig *interface{} `json:"cdm_params_config,omitempty"`

	HostConfig *interface{} `json:"host_config,omitempty"`
}

func (o AuthConfigA) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthConfigA struct{}"
	}

	return strings.Join([]string{"AuthConfigA", string(data)}, " ")
}
