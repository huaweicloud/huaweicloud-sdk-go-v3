package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssistAuthConfigResponse Response Object
type ShowAssistAuthConfigResponse struct {
	OtpConfigInfo *OtpConfigInfo `json:"otp_config_info,omitempty"`

	// 主认证配置id
	MainAuthConfigId *string `json:"main_auth_config_id,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowAssistAuthConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssistAuthConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAssistAuthConfigResponse", string(data)}, " ")
}
