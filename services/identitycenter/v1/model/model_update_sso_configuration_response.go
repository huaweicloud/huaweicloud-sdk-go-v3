package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSsoConfigurationResponse Response Object
type UpdateSsoConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSsoConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSsoConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateSsoConfigurationResponse", string(data)}, " ")
}
