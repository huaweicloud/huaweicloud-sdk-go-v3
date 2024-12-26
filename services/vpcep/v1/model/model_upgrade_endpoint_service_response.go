package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeEndpointServiceResponse Response Object
type UpgradeEndpointServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpgradeEndpointServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"UpgradeEndpointServiceResponse", string(data)}, " ")
}
