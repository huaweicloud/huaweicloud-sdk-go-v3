package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRemediationConfigurationResponse Response Object
type DeleteRemediationConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRemediationConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRemediationConfigurationResponse struct{}"
	}

	return strings.Join([]string{"DeleteRemediationConfigurationResponse", string(data)}, " ")
}
