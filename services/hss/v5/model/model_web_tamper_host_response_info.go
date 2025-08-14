package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperHostResponseInfo 网页防篡改可选服务器信息
type WebTamperHostResponseInfo struct {

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: Agent ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： - Linux：Linux。 - Windows：Windows。
	OsType *string `json:"os_type,omitempty"`

	// 资产重要性，包含如下3种   - important ：重要资产   - common ：一般资产   - test ：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// web应用列表
	WebAppList *[]string `json:"web_app_list,omitempty"`
}

func (o WebTamperHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperHostResponseInfo struct{}"
	}

	return strings.Join([]string{"WebTamperHostResponseInfo", string(data)}, " ")
}
