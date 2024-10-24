package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceAccessControlAttributeConfigurationReqBody the request body of UpdateInstanceAccessControlAttributeConfiguration
type UpdateInstanceAccessControlAttributeConfigurationReqBody struct {
	InstanceAccessControlAttributeConfiguration *InstanceAccessControlAttributeConfigurationDto `json:"instance_access_control_attribute_configuration"`
}

func (o UpdateInstanceAccessControlAttributeConfigurationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceAccessControlAttributeConfigurationReqBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceAccessControlAttributeConfigurationReqBody", string(data)}, " ")
}
