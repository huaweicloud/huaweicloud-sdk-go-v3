package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyLoginWhiteIpRequestBody 设置白名单列表
type ModifyLoginWhiteIpRequestBody struct {

	// **参数解释**： 白名单启用状态 **约束限制**： 不涉及 **取值范围**： - true：已启用 - false：已禁用  **默认取值**： false
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 白名单IP或IP网段，IP网段由IP地址和掩码组成，以“/”连接，单一账号最多可添加10个SSH登录IP白名单。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	WhiteIp string `json:"white_ip"`

	// **参数解释**： 服务器列表，不可为NULL，删除白名单IP或IP网段时，需要将服务器ID列表置为空列表[]。 **约束限制**： 不可为NULL **取值范围**： 不涉及 **默认取值**： 不涉及
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o ModifyLoginWhiteIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyLoginWhiteIpRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyLoginWhiteIpRequestBody", string(data)}, " ")
}
