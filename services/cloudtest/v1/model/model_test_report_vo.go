package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestReportVo struct {

	// 测试报告Uri
	Uri *string `json:"uri,omitempty"`

	// 测试报告名称
	Name *string `json:"name,omitempty"`

	// 创建人ID
	Creator *string `json:"creator,omitempty"`

	// 修改人ID
	Updator *string `json:"updator,omitempty"`

	// 测试计划Uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 分支Uri
	BranchUri *string `json:"branch_uri,omitempty"`

	// 测试计划名称
	VersionName *string `json:"version_name,omitempty"`

	// 分支名称
	BranchName *string `json:"branch_name,omitempty"`

	// 测试结论
	TestConclusion *string `json:"test_conclusion,omitempty"`

	// 测试结论描述
	TestConclusionDetails *string `json:"test_conclusion_details,omitempty"`

	// 缺陷解决率
	DefectResolutionRate *string `json:"defect_resolution_rate,omitempty"`

	// 缺陷解决分数
	DefectResolutionScore *string `json:"defect_resolution_score,omitempty"`

	// 用例执行率
	CaseExecutionRate *string `json:"case_execution_rate,omitempty"`

	// 用例执行分数
	CaseExecutionScore *string `json:"case_execution_score,omitempty"`

	// 用例通过率
	CasePassRate *string `json:"case_pass_rate,omitempty"`

	// 用例通过分数
	CasePassScore *string `json:"case_pass_score,omitempty"`

	// 需求通过率
	IssuePassRate *string `json:"issue_pass_rate,omitempty"`

	// 需求通过分数
	IssuePassScore *string `json:"issue_pass_score,omitempty"`

	// 需求覆盖率
	IssueCoverageRate *string `json:"issue_coverage_rate,omitempty"`

	// 需求覆盖分数
	IssueCoverageScore *string `json:"issue_coverage_score,omitempty"`

	// 项目总遗留DI
	ProjectResidualDefectIndex *string `json:"project_residual_defect_index,omitempty"`

	// 计划新增DI
	IteratorResidualDefectIndex *string `json:"iterator_residual_defect_index,omitempty"`

	CaseAutomationDetails *CaseAutomationDetailsVo `json:"case_automation_details,omitempty"`

	// 用例有效性比例
	CaseValidityRatio *string `json:"case_validity_ratio,omitempty"`

	IssueDetails *IssuePassDetailsVo `json:"issue_details,omitempty"`

	// 用例通过情况
	CaseDetails *[]NameAndValueVo `json:"case_details,omitempty"`

	// 缺陷严重程度
	DefectDetailsBySeverity *[]IdAndNameAndValueVo `json:"defect_details_by_severity,omitempty"`

	// 缺陷按照模块分布情况
	DefectDetailsByModule *[]IdAndNameAndValueVo `json:"defect_details_by_module,omitempty"`

	// 每个测试类型的用例通过率
	CasePassRateByTestType *[]DetailTestTypeCasePassRateVo `json:"case_pass_rate_by_test_type,omitempty"`

	// 测试报告自定义报告详情
	TestReportCustomReportDetail *[]CustomReportListVo `json:"test_report_custom_report_detail,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 创建时间戳
	CreateTimestamp *int64 `json:"create_timestamp,omitempty"`

	// 创建人名
	CreatorName *string `json:"creator_name,omitempty"`

	// 修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 修改时间戳
	UpdateTimestamp *int64 `json:"update_timestamp,omitempty"`

	// 修改人名
	UpdatorName *string `json:"updator_name,omitempty"`

	// 项目ID
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 风险分析
	RiskAnalysis *string `json:"risk_analysis,omitempty"`
}

func (o TestReportVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestReportVo struct{}"
	}

	return strings.Join([]string{"TestReportVo", string(data)}, " ")
}
