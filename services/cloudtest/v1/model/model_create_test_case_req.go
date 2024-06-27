package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTestCaseReq struct {

	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`

	// 是否添加到计划
	AddToPlan *bool `json:"addToPlan,omitempty"`

	// 创建时可选择导入的aw目录直接产生测试步骤
	AwCataList *[]BasicAwCata `json:"aw_cata_list,omitempty"`

	AwInstance *CaseAwInstance `json:"aw_instance,omitempty"`

	// 用例类型：0商用现有类型，1从内部导过来的用例类型
	CaseType *int32 `json:"case_type,omitempty"`

	// 用例局部变量
	CaseVariableList *[]AwVariable `json:"case_variable_list,omitempty"`

	// 演示标志
	DemoFlag *string `json:"demoFlag,omitempty"`

	// 错误测试阶段
	ErrorStep *[]int32 `json:"errorStep,omitempty"`

	// id
	Id *string `json:"id,omitempty"`

	// import信息List
	ImportInfoList *[]string `json:"import_info_list,omitempty"`

	// 是否设置为关键字操作
	IsKeyWord *int32 `json:"isKeyWord,omitempty"`

	// 是否同步
	IsSync *bool `json:"isSync,omitempty"`

	// 问题ID
	IssueId *string `json:"issueId,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 被复制用例的tmsscaseuri
	OldTmssCaseUri *string `json:"old_tmss_case_uri,omitempty"`

	// 被复制用例的tmsscaseuri列表，内部使用
	OldTmssCaseUriList *[]string `json:"old_tmss_case_uri_list,omitempty"`

	// 包名
	PackageName *string `json:"package_name,omitempty"`

	// 计划ID
	PlanId *string `json:"planId,omitempty"`

	// 工程ID
	ProjectId *string `json:"project_id,omitempty"`

	// 新服务新建用例版本号
	ReleaseDev *string `json:"releaseDev,omitempty"`

	// 脚本名(类名)
	ScriptName *string `json:"script_name,omitempty"`

	// 脚本路径
	ScriptPath *string `json:"script_path,omitempty"`

	// 来源
	Source *string `json:"source,omitempty"`

	// 来源的方式
	SourceWay *string `json:"sourceWay,omitempty"`

	// 选择用例中测试步骤生成关键字时的原用例tmssCaseUri
	SrcTmssCaseUri *string `json:"src_tmss_case_uri,omitempty"`

	// tmss用例uri
	TmssCaseUri *string `json:"tmss_case_uri,omitempty"`

	// tmss用例uri
	TmssFeatureUri *string `json:"tmss_feature_uri,omitempty"`

	TmssProperty *TmssTestcase `json:"tmss_property,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 环境参数分组id
	VariableGroupId *string `json:"variableGroupId,omitempty"`
}

func (o CreateTestCaseReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTestCaseReq struct{}"
	}

	return strings.Join([]string{"CreateTestCaseReq", string(data)}, " ")
}
