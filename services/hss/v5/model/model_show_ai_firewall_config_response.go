package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiFirewallConfigResponse Response Object
type ShowAiFirewallConfigResponse struct {

	// **参数解释**: 防火墙AK（Access Key），用于身份标识 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	SourceId *string `json:"source_id,omitempty"`

	// **参数解释**: 防火墙SK（Secret Key），用于签名认证，解密返回 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	SourceSecret *string `json:"source_secret,omitempty"`

	// **参数解释**: 应用ID，标识接入防火墙的应用 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAiFirewallConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiFirewallConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAiFirewallConfigResponse", string(data)}, " ")
}
