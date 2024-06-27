package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskAssignCaseDetailVo 实际的数据类型：单个对象，集合 或 NULL
type TaskAssignCaseDetailVo struct {

	// 关联关系uri
	Uri *string `json:"uri,omitempty"`

	// 排序顺序
	Sort *int32 `json:"sort,omitempty"`

	// 责任人id
	Owner *string `json:"owner,omitempty"`

	// 测试阶段
	Stage *int32 `json:"stage,omitempty"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 任务uri
	TaskUri *string `json:"task_uri,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 更新人
	UpdatorName *string `json:"updator_name,omitempty"`

	// 更新人id
	Updator *string `json:"updator,omitempty"`

	// 用例uri
	CaseUri *string `json:"case_uri,omitempty"`

	// 是否可用
	IsAvailable *int32 `json:"is_available,omitempty"`

	// 用例名称
	TestCaseName *string `json:"test_case_name,omitempty"`

	// 用例目录Uri
	FeatureUri *string `json:"feature_uri,omitempty"`

	// 用例编号
	TestCaseNumber *string `json:"test_case_number,omitempty"`

	// 脚本路径
	SvnScriptPath *string `json:"svn_script_path,omitempty"`

	// 状态
	StatusCode *string `json:"status_code,omitempty"`

	// 状态名称
	StatusName *string `json:"status_name,omitempty"`

	// 结果id
	ResultCode *string `json:"result_code,omitempty"`

	// 结果名称
	ResultName *string `json:"result_name,omitempty"`

	// 责任人名称
	OwnerName *string `json:"owner_name,omitempty"`

	// 最新执行时间
	ExecuteLatestTime *sdktime.SdkTime `json:"execute_latest_time,omitempty"`

	// 执行时长
	ExecuteDuration *string `json:"execute_duration,omitempty"`

	// 是否是关键用例
	IsKeyword *int32 `json:"is_keyword,omitempty"`

	// 脚本名称
	NetWorkScriptName *string `json:"net_work_script_name,omitempty"`

	// 用例等级
	RankId *int32 `json:"rank_id,omitempty"`
}

func (o TaskAssignCaseDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAssignCaseDetailVo struct{}"
	}

	return strings.Join([]string{"TaskAssignCaseDetailVo", string(data)}, " ")
}
