package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVrrpConfigResponse Response Object
type ShowVrrpConfigResponse struct {
	VrrpConfigs    *[]VrrpConfigItem `json:"vrrp_configs,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowVrrpConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVrrpConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowVrrpConfigResponse", string(data)}, " ")
}
