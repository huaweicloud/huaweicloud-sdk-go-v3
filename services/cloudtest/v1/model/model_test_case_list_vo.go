package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestCaseListVo 实际的数据类型：单个对象，集合 或 NULL
type TestCaseListVo struct {

	// 用例URI
	Uri *string `json:"uri,omitempty"`

	// 用例名称
	Name *string `json:"name,omitempty"`

	Owner *NameAndIdVo `json:"owner,omitempty"`

	Status *NameAndIdVo `json:"status,omitempty"`

	Result *NameAndIdVo `json:"result,omitempty"`

	Module *NameAndIdVo `json:"module,omitempty"`

	Iteration *NameAndIdVo `json:"iteration,omitempty"`

	// 执行平台，如：apittest,其他三方执行平台
	Exeplatform *string `json:"exeplatform,omitempty"`

	// 用例编号
	Number *string `json:"number,omitempty"`

	// 用例描述
	Description *string `json:"description,omitempty"`

	// 用例等级
	RankId *string `json:"rank_id,omitempty"`

	// 目录URI
	FeatureUri *string `json:"feature_uri,omitempty"`

	// 版本号
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 是否组合关键字
	IsKeyword *bool `json:"is_keyword,omitempty"`

	// 脚本路径
	ScriptUrl *string `json:"script_url,omitempty"`

	// 实时报告地址
	ReportUrl *string `json:"report_url,omitempty"`

	// 项目ID
	ProjectUuid *string `json:"project_uuid,omitempty"`

	ServiceType *NameAndIdVo `json:"service_type,omitempty"`

	TestType *IntegerIdAndNameVo `json:"test_type,omitempty"`

	CreateInfo *CreateInfoVo `json:"create_info,omitempty"`

	ExecuteInfo *ExecuteInfoVo `json:"execute_info,omitempty"`

	AssociateIssueInfo *AssociateIssueInfoVo `json:"associate_issue_info,omitempty"`

	AssociateDefectInfo *AssociateDefectInfoVo `json:"associate_defect_info,omitempty"`

	// 用例类型
	CaseType *int32 `json:"case_type,omitempty"`

	// 用例标签名称列表
	Labels *string `json:"labels,omitempty"`

	// 自定义字段信息
	CustomFieldInfo *[]CustomFieldVo `json:"custom_field_info,omitempty"`

	// 是否来自测试设计（null：不限，false：否来自测试设计，true：来自测试设计）
	IsTestDesign *bool `json:"is_test_design,omitempty"`

	// 最后修改时间（null：不限）
	LastModified *int64 `json:"last_modified,omitempty"`

	// 用例评审状态（null：0至127）
	ReviewStatus *int32 `json:"review_status,omitempty"`
}

func (o TestCaseListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCaseListVo struct{}"
	}

	return strings.Join([]string{"TestCaseListVo", string(data)}, " ")
}
