package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlanResponseResult **参数解释**： 响应中的计划详情数据对象，日期字段为unix时间戳格式。 **约束限制**： 不涉及。
type PlanResponseResult struct {

	// **参数解释：** 发布、迭代、里程碑的ID **取值范围：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 标题 **取值范围：** 不涉及
	Title *string `json:"title,omitempty"`

	// **参数解释：** 分类，枚举类型 **取值范围：** - PI：发布 - Iteration：迭代 - PlanMilestone：里程碑
	Category *string `json:"category,omitempty"`

	// **参数解释：** 描述 **取值范围：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 作废标识，枚举类型。 **取值范围：** - 正在工作：可正常操作的发布。 - 作废：软删除后的发布。 - 删除：彻底删除后的发布。
	State *string `json:"state,omitempty"`

	// **参数解释：** 发布/迭代的状态，枚举类型。 **取值范围：** - planned：发布/计划未开始 - going：发布/计划进行中 - ended：发布/计划已结束
	Status *string `json:"status,omitempty"`

	// **参数解释：** 子项目迭代信息
	Children *[]PlanResponseResult `json:"children,omitempty"`

	// **参数解释：** 创建人ID **取值范围：** 不涉及
	CreatedBy *string `json:"created_by,omitempty"`

	// **参数解释：** 最近更新人ID。 **取值范围：** 不涉及
	ModifiedBy *string `json:"modified_by,omitempty"`

	// **参数解释：** 计划开始时间，unix时间戳，单位：毫秒。 **取值范围：** 不涉及
	PlanStartDate *int64 `json:"plan_start_date,omitempty"`

	// **参数解释：** 计划完成时间，unix时间戳，单位：毫秒。 **取值范围：** 不涉及
	PlanEndDate *int64 `json:"plan_end_date,omitempty"`

	// **参数解释：** 创建时间，unix时间戳，单位：毫秒 **取值范围：** 不涉及
	CreatedDate *int64 `json:"created_date,omitempty"`

	// **参数解释：** 父计划ID，当计划类型为迭代(Iteration)时，用于指定所属的发布计划。 **取值范围：** 长度为18~19个字符的数字字符串。
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释：** 基线状态，枚举类型，标识计划是否已纳入基线管理。 **取值范围：** - baselined：已基线 - unbaseline：未基线 - \"\"：未基线
	Baseline *string `json:"baseline,omitempty"`

	// **参数解释：** 预估工作量，用于标识计划所需的人力或时间投入，单位人/天。 **取值范围：** 最大长度11个字符。
	Workload *string `json:"workload,omitempty"`

	// **参数解释：** 责任人ID，标识计划的负责人。 **取值范围：** 长度为32位的字符串。
	Owner *string `json:"owner,omitempty"`
}

func (o PlanResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlanResponseResult struct{}"
	}

	return strings.Join([]string{"PlanResponseResult", string(data)}, " ")
}
