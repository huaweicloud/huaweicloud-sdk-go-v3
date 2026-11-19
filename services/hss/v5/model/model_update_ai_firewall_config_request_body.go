package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiFirewallConfigRequestBody 配置大模型防火墙请求结构体
type UpdateAiFirewallConfigRequestBody struct {

	// **参数解释**: 防火墙AK（Access Key），用于身份标识，明文传入，加密存储 **约束限制**: 非必填，但配置防火墙检测时需要与source_secret、app_id同时配置 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	SourceId *string `json:"source_id,omitempty"`

	// **参数解释**: 防火墙SK（Secret Key），用于签名认证，明文传入，加密存储 **约束限制**: 非必填，但配置防火墙检测时需要与source_id、app_id同时配置 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	SourceSecret *string `json:"source_secret,omitempty"`

	// **参数解释**: 应用ID，标识接入防火墙的应用 **约束限制**: 非必填，但配置防火墙检测时需要与source_id、source_secret同时配置 **取值范围**: 字符长度0-64位 **默认取值**: 不涉及
	AppId *string `json:"app_id,omitempty"`
}

func (o UpdateAiFirewallConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiFirewallConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAiFirewallConfigRequestBody", string(data)}, " ")
}
