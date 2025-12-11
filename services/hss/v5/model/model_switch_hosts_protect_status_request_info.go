package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchHostsProtectStatusRequestInfo 切换防护的请求信息
type SwitchHostsProtectStatusRequestInfo struct {

	// **参数解释**： 主机开通的版本 **约束限制**: 不涉及 **取值范围**： 包含如下7种输入。 - hss.version.null ：无。 - hss.version.basic ：基础版。 - hss.version.advanced ：专业版。 - hss.version.enterprise ：企业版。 - hss.version.premium ：旗舰版。 - hss.version.wtp ：网页防篡改版。 - hss.version.container.enterprise：容器版。 **默认取值**: 不涉及
	Version string `json:"version"`

	// **参数解释**： 计费模式 **约束限制**: 不涉及 **取值范围**： - packet_cycle ：包年/包月。 - on_demand ：按需计费。 **默认取值**: 不涉及
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释**: 资源ID **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 服务器ID列表 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
	HostIdList []string `json:"host_id_list"`

	// **参数解释**： 资源标签列表 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
	Tags *[]TagInfo `json:"tags,omitempty"`
}

func (o SwitchHostsProtectStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchHostsProtectStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"SwitchHostsProtectStatusRequestInfo", string(data)}, " ")
}
