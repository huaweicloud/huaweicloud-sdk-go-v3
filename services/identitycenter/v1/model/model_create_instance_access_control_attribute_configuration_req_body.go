package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceAccessControlAttributeConfigurationReqBody the request body of CreateInstanceAccessControlAttributeConfiguration
type CreateInstanceAccessControlAttributeConfigurationReqBody struct {
	InstanceAccessControlAttributeConfiguration *InstanceAccessControlAttributeConfigurationDto `json:"instance_access_control_attribute_configuration"`
}

func (o CreateInstanceAccessControlAttributeConfigurationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceAccessControlAttributeConfigurationReqBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceAccessControlAttributeConfigurationReqBody", string(data)}, " ")
}
