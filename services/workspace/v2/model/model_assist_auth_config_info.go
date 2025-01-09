package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssistAuthConfigInfo 辅助认证配置信息。
type AssistAuthConfigInfo struct {

	// 当前激活的辅助认证类型。
	AuthType *string `json:"auth_type,omitempty"`

	OtpConfigInfo *OtpConfigInfo `json:"otp_config_info,omitempty"`
}

func (o AssistAuthConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssistAuthConfigInfo struct{}"
	}

	return strings.Join([]string{"AssistAuthConfigInfo", string(data)}, " ")
}
