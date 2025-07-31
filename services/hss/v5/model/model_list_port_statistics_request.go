package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortStatisticsRequest Request Object
type ListPortStatisticsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 类别 **约束限制**: 不涉及 **取值范围**: - host：主机 - container：容器  **默认取值**: 不涉及
	Category *string `json:"category,omitempty"`

	// **参数解释**: 排序的顺序 **约束限制**: 不涉及 **取值范围**:   - asc  : 正序   - desc : 倒序  **默认取值**: 正序排序
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 端口号，该字段用来进行精确匹配 **约束限制**: 与port_string同时使用的话，二者有包含关系则按精确匹配，无包含关系则结果 为空 **取值范围**: 最小值1，最大值65535 **默认取值**: 不涉及
	Port *int32 `json:"port,omitempty"`

	// **参数解释**: 端口字符串，该字段用来进行模糊匹配 **约束限制**: 与port同时使用的话，二者有包含关系则按精确匹配，无包含关系则结果 为空 **取值范围**: 最小值1，最大值65535 **默认取值**: 不涉及
	PortString *string `json:"port_string,omitempty"`

	// **参数解释**: 端口类型 **约束限制**: 不涉及 **取值范围**: - UDP - UDP6 - TCP - TCP6  **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**: 端口状态 **约束限制**: 不涉及 **取值范围**: - danger: 危险端口 - normal: 正常端口 - unknow: 无已知危险的端口  **默认取值**: 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**: 排序的key值，目前支持按照端口号port排序 **约束限制**: 不涉及 **取值范围**: - port: 端口号  **默认取值**: 不涉及
	SortKey *string `json:"sort_key,omitempty"`
}

func (o ListPortStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListPortStatisticsRequest", string(data)}, " ")
}
