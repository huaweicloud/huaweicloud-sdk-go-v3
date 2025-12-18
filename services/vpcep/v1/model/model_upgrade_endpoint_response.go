package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeEndpointResponse Response Object
type UpgradeEndpointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpgradeEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeEndpointResponse struct{}"
	}

	return strings.Join([]string{"UpgradeEndpointResponse", string(data)}, " ")
}
