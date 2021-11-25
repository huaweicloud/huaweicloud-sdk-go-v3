package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdatePremiumHostProtectStatusResponse struct {
	// 域名防护状态：  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测

	ProtectStatus  *int32 `json:"protect_status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdatePremiumHostProtectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostProtectStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostProtectStatusResponse", string(data)}, " ")
}
