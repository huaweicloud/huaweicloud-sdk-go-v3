package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConfigurationResponse Response Object
type UpdateConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationResponse", string(data)}, " ")
}
