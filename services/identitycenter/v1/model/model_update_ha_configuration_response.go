package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHaConfigurationResponse Response Object
type UpdateHaConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateHaConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHaConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateHaConfigurationResponse", string(data)}, " ")
}
