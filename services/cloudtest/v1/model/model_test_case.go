package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestCase struct {

	// 自动化类型
	AutoType *string `json:"auto_type,omitempty"`

	// 用例设计描述
	CaseDesignDesc *string `json:"case_design_desc,omitempty"`

	// 用例名称
	CaseName *string `json:"case_name,omitempty"`

	// 用例编号
	CaseNum *string `json:"case_num,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 创建人名字
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建人工号
	CreatorNum *string `json:"creator_num,omitempty"`

	// 删除状态
	Deleted *string `json:"deleted,omitempty"`

	// 预期结果
	ExpectedResults *string `json:"expected_results,omitempty"`

	// 额外模板字段：以json形式存储，前台解析
	ExtraParam *string `json:"extra_param,omitempty"`

	// 因子组合json
	FactorCombinationJson *string `json:"factor_combination_json,omitempty"`

	// 操作及预期结果
	OperationAndExpectedResult *string `json:"operation_and_expected_result,omitempty"`

	// id 主键
	Id *string `json:"id,omitempty"`

	// 状态
	IsArchive *bool `json:"is_archive,omitempty"`

	// 脑图id
	MindmapId *string `json:"mindmap_id,omitempty"`

	// 节点id
	NodeId *string `json:"node_id,omitempty"`

	// 批次id
	BatchId *string `json:"batch_id,omitempty"`

	// 分支ID
	BranchId *string `json:"branch_id,omitempty"`

	// 计划ID
	PlanId *string `json:"plan_id,omitempty"`

	// 用例前置步骤
	Prerequisite *string `json:"prerequisite,omitempty"`

	// 测试用例级别
	TestCaseLevel *string `json:"test_case_level,omitempty"`

	// 测试步骤
	TestProcedure *string `json:"test_procedure,omitempty"`

	// 更新人名字
	UpdateName *string `json:"update_name,omitempty"`

	// 更新人工号
	UpdateNum *string `json:"update_num,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// url
	Url *string `json:"url,omitempty"`

	// uri
	Uri *string `json:"uri,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 服务id
	ServiceId *string `json:"service_id,omitempty"`
}

func (o TestCase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCase struct{}"
	}

	return strings.Join([]string{"TestCase", string(data)}, " ")
}
