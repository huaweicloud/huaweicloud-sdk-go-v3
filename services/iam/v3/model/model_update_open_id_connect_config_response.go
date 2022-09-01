package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateOpenIdConnectConfigResponse struct {
	OpenidConnectConfig *OpenIdConnectConfig `json:"openid_connect_config,omitempty" xml:"openid_connect_config"`
	HttpStatusCode      int                  `json:"-"`
}

func (o UpdateOpenIdConnectConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOpenIdConnectConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateOpenIdConnectConfigResponse", string(data)}, " ")
}
