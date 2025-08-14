package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceSecurityConfigurationReqBody UpdateApplicationInstanceSecurityConfiguration的请求体
type UpdateApplicationInstanceSecurityConfigurationReqBody struct {
	SecurityConfig *SecurityConfigDto `json:"security_config"`
}

func (o UpdateApplicationInstanceSecurityConfigurationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceSecurityConfigurationReqBody struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceSecurityConfigurationReqBody", string(data)}, " ")
}
