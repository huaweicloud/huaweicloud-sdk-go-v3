package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceAccessControlAttributeConfigurationResponse Response Object
type UpdateInstanceAccessControlAttributeConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceAccessControlAttributeConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceAccessControlAttributeConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceAccessControlAttributeConfigurationResponse", string(data)}, " ")
}
