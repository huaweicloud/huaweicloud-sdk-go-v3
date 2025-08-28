package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHandleVulsRequest Request Object
type ListHandleVulsRequest struct {

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 漏洞名称 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	VulName *string `json:"vul_name,omitempty"`

	// **参数解释**: 漏洞cve编号 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	CveId *string `json:"cve_id,omitempty"`

	// **参数解释**: 漏洞标签 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	LabelList *string `json:"label_list,omitempty"`

	// **参数解释**： 漏洞状态 **约束限制**： 不涉及 **取值范围**： - vul_status_ignored: 已忽略。 - vul_status_fixed: 修复成功。  **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**： 资产重要性 **约束限制**: 不涉及 **取值范围**: - important：重要资产。 - common：一般资产。 - test：测试资产。  **默认取值**: 不涉及
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**: 服务器组名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 主机名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机ip **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 漏洞类型 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 处理周期 **约束限制**: 不涉及 **取值范围**: - today：今日处理。 - all：累计处理。  **默认取值**: 不涉及
	HandleCircle *string `json:"handle_circle,omitempty"`

	// **参数解释**： 修复优先级 **约束限制**: 不涉及 **取值范围**: - Critical：紧急。 - High：高。 - Medium：中。 - Low：低。  **默认取值**: 不涉及
	RepairPriority *string `json:"repair_priority,omitempty"`
}

func (o ListHandleVulsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHandleVulsRequest struct{}"
	}

	return strings.Join([]string{"ListHandleVulsRequest", string(data)}, " ")
}
