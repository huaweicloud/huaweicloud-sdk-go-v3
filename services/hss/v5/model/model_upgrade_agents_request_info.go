package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradeAgentsRequestInfo struct {

	// **参数解释**: 是否全量升级 **约束限制**: 不涉及 **取值范围**: - true: 全量升级，不需要填写host_id_list，按照其余筛选条件筛选符合升级的服务器。 - false: 非全量升级，需要填写host_id_list，只升级host_id_list中包含的服务器。  **默认取值**: false
	OperateAll bool `json:"operate_all"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 弹性公网IP地址 **取值范围**： 字符长度1-256位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 防护版本 **约束限制**: 不涉及 **取值范围**： - hss.version.basic ：基础版。 - hss.version.advanced ：专业版。 - hss.version.enterprise ：企业版。 - hss.version.premium ：旗舰版。 - hss.version.wtp ：网页防篡改版。 - hss.version.container.enterprise：容器版。  **默认取值**: 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**： 防护状态 **约束限制**: 不涉及 **取值范围**： - closed ：未防护。 - opened ：防护中。  **默认取值**: 不涉及
	ProtectStatus *string `json:"protect_status,omitempty"`

	// **参数解释**： 操作系统类型 **约束限制**: 不涉及 **取值范围**： - Linux：Linux系统。 - Windows：Windows系统。  **默认取值**: 不涉及
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 策略组ID **取值范围**： 字符长度36-64位
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// **参数解释**: 服务器组ID **取值范围**: 字符长度0-64位
	GroupId *string `json:"group_id,omitempty"`

	// 资产重要性 **参数解释**： 资产重要性 **约束限制**: 不涉及 **取值范围**： - important ：重要资产。 - common ：一般资产。 - test ：测试资产。  **默认取值**: 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 服务器ID列表 **取值范围**: 不涉及
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o UpgradeAgentsRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeAgentsRequestInfo struct{}"
	}

	return strings.Join([]string{"UpgradeAgentsRequestInfo", string(data)}, " ")
}
