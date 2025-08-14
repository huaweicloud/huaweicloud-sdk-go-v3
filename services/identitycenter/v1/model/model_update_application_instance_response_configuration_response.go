package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceResponseConfigurationResponse Response Object
type UpdateApplicationInstanceResponseConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateApplicationInstanceResponseConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceResponseConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceResponseConfigurationResponse", string(data)}, " ")
}
