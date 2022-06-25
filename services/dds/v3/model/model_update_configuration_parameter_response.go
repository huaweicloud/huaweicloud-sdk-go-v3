package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateConfigurationParameterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConfigurationParameterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigurationParameterResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationParameterResponse", string(data)}, " ")
}
