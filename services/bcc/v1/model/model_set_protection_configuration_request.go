package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetProtectionConfigurationRequest Request Object
type SetProtectionConfigurationRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	Body *ConfigureProtectionBody `json:"body,omitempty"`
}

func (o SetProtectionConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetProtectionConfigurationRequest struct{}"
	}

	return strings.Join([]string{"SetProtectionConfigurationRequest", string(data)}, " ")
}
