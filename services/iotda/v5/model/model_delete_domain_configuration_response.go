package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainConfigurationResponse Response Object
type DeleteDomainConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDomainConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainConfigurationResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainConfigurationResponse", string(data)}, " ")
}
