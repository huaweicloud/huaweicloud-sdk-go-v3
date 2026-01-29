package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PlanVo struct {

	// **参数解释：** id（发布、迭代、里程碑的id） **取值范围：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 标题 **取值范围：** 不涉及
	Title *string `json:"title,omitempty"`

	// **参数解释：** 分类，枚举类型 **取值范围：** - PI：发布 - Iteration：迭代 - PlanMilestone：里程碑
	Category *string `json:"category,omitempty"`

	// **参数解释：** 描述 **取值范围：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 作废标识，枚举类型。 **取值范围：** - 正在工作：可正常操作的发布 - 作废：软删除后的发布，可在回收站恢复 - 删除：彻底删除后的发布，无法恢复
	State *string `json:"state,omitempty"`

	// **参数解释：** 发布/迭代的状态，枚举类型。 **取值范围：** - planned：发布/计划未开始 - going：发布/计划进行中 - ended：发布/计划已结束
	Status *string `json:"status,omitempty"`

	// **参数解释：** 子项目迭代信息
	Children *[]PlanVo `json:"children,omitempty"`

	// **参数解释：** 创建人Id **取值范围：** 不涉及
	CreatedBy *string `json:"created_by,omitempty"`

	// **参数解释：** 最近更新人Id **取值范围：** 不涉及
	ModifiedBy *string `json:"modified_by,omitempty"`

	// **参数解释：** 计划开始时间，\"yyyy-MM-dd\"格式 **取值范围：** 不涉及
	PlanStartDate *string `json:"plan_start_date,omitempty"`

	// **参数解释：** 计划完成时间，\"yyyy-MM-dd\"格式 **取值范围：** 不涉及
	PlanEndDate *string `json:"plan_end_date,omitempty"`

	// **参数解释：** 创建时间，unix时间戳，单位：毫秒 **取值范围：** 不涉及
	CreatedDate *int64 `json:"created_date,omitempty"`

	// **参数解释：** 父工作项id **取值范围：** 不涉及
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释：** 基线状态，枚举类型 **取值范围：** - baselined：已基线 - unbaseline：未基线 - baseline-reviewing：基线评审中
	Baseline *string `json:"baseline,omitempty"`

	// **参数解释：** 预估工作量 **取值范围：** 不涉及
	Workload *string `json:"workload,omitempty"`

	// **参数解释：** 责任人ID **取值范围：** 不涉及
	Owner *string `json:"owner,omitempty"`
}

func (o PlanVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlanVo struct{}"
	}

	return strings.Join([]string{"PlanVo", string(data)}, " ")
}
