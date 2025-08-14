package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceAccessControlAttributeConfigurationDto Specifies the attributes to add to your attribute-based access control (ABAC) configuration.
type InstanceAccessControlAttributeConfigurationDto struct {

	// IAM Identity Center实例中ABAC配置的属性
	AccessControlAttributes []AccessControlAttributeDto `json:"access_control_attributes"`
}

func (o InstanceAccessControlAttributeConfigurationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceAccessControlAttributeConfigurationDto struct{}"
	}

	return strings.Join([]string{"InstanceAccessControlAttributeConfigurationDto", string(data)}, " ")
}
