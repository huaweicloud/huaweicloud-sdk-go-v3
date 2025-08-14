package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceSecurityConfigurationResponse Response Object
type UpdateApplicationInstanceSecurityConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateApplicationInstanceSecurityConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceSecurityConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceSecurityConfigurationResponse", string(data)}, " ")
}
