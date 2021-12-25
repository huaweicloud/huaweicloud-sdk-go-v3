package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowPlanJournalsResponseDetail struct {
	// 测试计划基础信息变更，包括计划名称，测试类型，计划处理者、版本号、关联迭代、开始日期、截至日期、描述

	Updated *[]AttributeChange `json:"updated,omitempty"`
	// 测试计划资源的添加记录（工作项或者测试用例）

	Added *[]NameAndId `json:"added,omitempty"`
	// 测试计划资源的移除记录（工作项或者测试用例）

	Deleted *[]NameAndId `json:"deleted,omitempty"`
	// 表明该条变更记录的具体变更类型，例如测试用例（testCase），需求（issue）

	JournalizedType *string `json:"journalized_type,omitempty"`
	// 表明该条变更记录属于基础信息变更还是资源（需求添加移除、用例添加移除）变更

	Type *string `json:"type,omitempty"`
}

func (o ShowPlanJournalsResponseDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlanJournalsResponseDetail struct{}"
	}

	return strings.Join([]string{"ShowPlanJournalsResponseDetail", string(data)}, " ")
}
