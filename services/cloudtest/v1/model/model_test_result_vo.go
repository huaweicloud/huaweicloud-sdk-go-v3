package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestResultVo 实际的数据类型：单个对象，集合 或 NULL
type TestResultVo struct {

	// 结果URI
	Uri *string `json:"uri,omitempty"`

	// 用例结果名称
	Name *string `json:"name,omitempty"`

	// 创建人ID
	Author *string `json:"author,omitempty"`

	// 级别
	Rank *int32 `json:"rank,omitempty"`

	// 测试结果Code
	Result *int32 `json:"result,omitempty"`

	// 执行批次
	Round *int32 `json:"round,omitempty"`

	// 前置条件
	Preparation *string `json:"preparation,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 逻辑Region
	Region *string `json:"region,omitempty"`

	// 测试步骤信息
	Steps *[]ResultStepVo `json:"steps,omitempty"`

	// 用例结果编号
	Number *string `json:"number,omitempty"`

	// 创建人名称
	AuthorName *string `json:"author_name,omitempty"`

	// 执行开始时间
	BeginTime *sdktime.SdkTime `json:"begin_time,omitempty"`

	// 执行开始时间时间戳
	BeginTimeTimestamp *int64 `json:"begin_time_timestamp,omitempty"`

	// 执行结束时间
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 执行结束时间时间戳
	EndTimeTimestamp *int64 `json:"end_time_timestamp,omitempty"`

	// 创建时间
	CreationDate *string `json:"creation_date,omitempty"`

	// 创建时间时间戳
	CreationDateTimestamp *int64 `json:"creation_date_timestamp,omitempty"`

	// 最后修改时间
	LastModified *sdktime.SdkTime `json:"last_modified,omitempty"`

	// 最后修改时间
	LastModifiedTimestamp *int64 `json:"last_modified_timestamp,omitempty"`

	// 最后变更时间
	LastChangeTime *sdktime.SdkTime `json:"last_change_time,omitempty"`

	// 最后变更时间
	LastChangeTimeTimestamp *int64 `json:"last_change_time_timestamp,omitempty"`

	// DFX测试结果
	DfxTestResult *string `json:"dfx_test_result,omitempty"`

	// 失败原因
	FailureCause *string `json:"failure_cause,omitempty"`

	// 父节点URI(分支用例URI或迭代用例URI)
	ParentUri *string `json:"parent_uri,omitempty"`

	// 分支用例URI
	TestCaseUri *string `json:"test_case_uri,omitempty"`

	// 用例名称
	TestCaseName *string `json:"test_case_name,omitempty"`

	// 测试任务URI
	TaskUri *string `json:"task_uri,omitempty"`

	// 测试结果
	ResultName *string `json:"result_name,omitempty"`

	// 是否自动化执行
	TestResultAe *string `json:"test_result_ae,omitempty"`

	// 执行人ID
	ExecutorId *string `json:"executor_id,omitempty"`

	// 执行人名称
	ExecutorName *string `json:"executor_name,omitempty"`

	// 执行机任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 执行ID
	ExecuteId *string `json:"execute_id,omitempty"`

	// 执行耗时
	TimeCost *int32 `json:"time_cost,omitempty"`

	// 测试步骤
	StepTxt *string `json:"step_txt,omitempty"`

	// 测试步骤期望结果
	StepExpect *string `json:"step_expect,omitempty"`

	// 测试步骤实际结果
	StepActual *string `json:"step_actual,omitempty"`

	// 测试步骤结果
	StepResult *string `json:"step_result,omitempty"`

	// 版本号
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 创建版本URI
	CreationVersionUri *string `json:"creation_version_uri,omitempty"`

	// 版本URI
	VersionUri *string `json:"version_uri,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 第三方过来的执行结果，返回跳转到第三方的url
	ReportUrl *string `json:"report_url,omitempty"`

	// 测试用例编号
	TestCaseNumber *string `json:"test_case_number,omitempty"`

	// 测试类型
	ServiceType *string `json:"service_type,omitempty"`
}

func (o TestResultVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestResultVo struct{}"
	}

	return strings.Join([]string{"TestResultVo", string(data)}, " ")
}
