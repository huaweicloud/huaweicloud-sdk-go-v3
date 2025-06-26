package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigurationResponse Response Object
type UpdateInstanceConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationResponse", string(data)}, " ")
}
