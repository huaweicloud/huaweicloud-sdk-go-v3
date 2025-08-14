package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotasDetailRequest Request Object
type ListQuotasDetailRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 主机开通的版本 **约束限制**: 不涉及 **取值范围**： 包含如下7种输入。 - hss.version.null ：无。 - hss.version.basic ：基础版。 - hss.version.advanced ：专业版。 - hss.version.enterprise ：企业版。 - hss.version.premium ：旗舰版。 - hss.version.wtp ：网页防篡改版。 - hss.version.container.enterprise：容器版。 **默认取值**: 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**: 类别 **约束限制**: 不涉及 **取值范围**: 包含如下两种： - host_resource ：主机 - container_resource ：容器 **默认取值**: 不涉及
	Category *string `json:"category,omitempty"`

	// **参数解释**: 配额状态 **约束限制**: 不涉及 **取值范围**: 包含如下三种： - normal ： 正常 - expired ：过期 - freeze ：冻结 **默认取值**: 不涉及
	QuotaStatus *string `json:"quota_status,omitempty"`

	// **参数解释**: 使用状态 **约束限制**: 不涉及 **取值范围**: 包含如下两种： - idle ：空闲的 - used ：使用中 **默认取值**: 不涉及
	UsedStatus *string `json:"used_status,omitempty"`

	// **参数解释**: 服务器名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-256位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** : HSS配额的资源ID **约束限制** : 不涉及 **取值范围** : 字符长度1-128位 **默认取值** : 不涉及
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 收费模式 **约束限制**: 不涉及 **取值范围**: - packet_cycle ：包年/包月。 - on_demand ：按需。 **默认取值**: 不涉及
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListQuotasDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasDetailRequest struct{}"
	}

	return strings.Join([]string{"ListQuotasDetailRequest", string(data)}, " ")
}
