package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWtpProtectHostRequest Request Object
type ListWtpProtectHostRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器弹性IP地址。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 无
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 服务器私有IP **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 服务器组名称。 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 无
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 操作系统类型。 **约束限制**: 不涉及 **取值范围**: - Linux：Linux操作系统。 - Windows：Windows操作系统。  **默认取值**: 无
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 资产重要性 **约束限制**： 不涉及 **取值范围**： - important：重要资产。 - common：一般资产。 - test：测试资产。  **默认取值**： 无
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 网页防篡改防护开启状态。 **约束限制**: 不涉及 **取值范围**: - opened ：已开启网页防篡改防护。  **默认取值**: 无
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**: 网页防篡改详细防护状态。 **约束限制**: 不涉及 **取值范围**: - opened : 防护中。 - opening : 开启中。 - open_failed : 防护失败。 - partial_protection : 部分防护。 - protection_interruption : 防护中断。  **默认取值**: 无
	WtpStatus *string `json:"wtp_status,omitempty"`

	// **参数解释**: Agent状态。 **约束限制**: 不涉及 **取值范围**: - not_installed : agent未安装。 - online : agent在线。 - offline : agent不在线。  **默认取值**: 无
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**: 动态网页防篡改防护开启状态。 **约束限制**: 不涉及 **取值范围**: - opened ：已开启动态网页防篡改防护。 - closed ：未开启动态网页防篡改防护。  **默认取值**: 无
	RaspStatus *string `json:"rasp_status,omitempty"`
}

func (o ListWtpProtectHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWtpProtectHostRequest struct{}"
	}

	return strings.Join([]string{"ListWtpProtectHostRequest", string(data)}, " ")
}
