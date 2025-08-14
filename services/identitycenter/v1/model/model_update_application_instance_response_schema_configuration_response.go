package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceResponseSchemaConfigurationResponse Response Object
type UpdateApplicationInstanceResponseSchemaConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateApplicationInstanceResponseSchemaConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceResponseSchemaConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceResponseSchemaConfigurationResponse", string(data)}, " ")
}
