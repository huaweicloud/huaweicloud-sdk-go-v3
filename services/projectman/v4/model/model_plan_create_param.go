package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlanCreateParam **参数解释**： 创建发布/迭代计划的请求参数。 **约束限制**： 不涉及。
type PlanCreateParam struct {

	// **参数解释**： 计划标题。 **约束限制**： 不涉及。 **取值范围**： 1~256个字符。 **默认取值**： 不涉及。
	Title string `json:"title"`

	// **参数解释**： 计划分类，枚举类型。 **约束限制**： 不涉及。 **取值范围**： - PI：发布 - Iteration：迭代 - PlanMilestone：里程碑 **默认取值**： 不涉及。
	Category string `json:"category"`

	// **参数解释**： 计划描述信息。 **约束限制**： 不涉及。 **取值范围**： 0~1000个字符。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 计划开始时间，格式为yyyy-MM-dd，如2024-01-01。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PlanStartDate string `json:"plan_start_date"`

	// **参数解释**： 计划完成时间，格式为yyyy-MM-dd，如2024-01-01。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PlanEndDate string `json:"plan_end_date"`

	// **参数解释**： 父计划ID，当category为Iteration时必填，用于指定所属的发布计划。 **约束限制**： category为Iteration时必填。 **取值范围**： 长度为18~19个字符的数字字符串。 **默认取值**： 不涉及。
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释**： 预估工作量，用于标识计划所需的人力投入，单位人/天。 **约束限制**： 不涉及。 **取值范围**： 0~11个字符。 **默认取值**： 不涉及。
	Workload *string `json:"workload,omitempty"`

	// **参数解释**： 责任人ID，标识计划的负责人。 **约束限制**： 不涉及。 **取值范围**： 长度为32个字符。 **默认取值**： 不涉及。
	Owner *string `json:"owner,omitempty"`
}

func (o PlanCreateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlanCreateParam struct{}"
	}

	return strings.Join([]string{"PlanCreateParam", string(data)}, " ")
}
