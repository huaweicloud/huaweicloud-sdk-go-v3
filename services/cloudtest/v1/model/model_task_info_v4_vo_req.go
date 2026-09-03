package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInfoV4VoReq struct {

	// 智能告警开关：0为置灰，1为可用
	AlertAction *string `json:"alert_action,omitempty"`

	// 告警配置
	AlertConfig *interface{} `json:"alert_config,omitempty"`

	// 流水线启动测试套件，携带构建产物
	BuildProducts *[]BuildProduct `json:"build_products,omitempty"`

	// 用例环境参数
	CasePackageEnvName *string `json:"case_package_env_name,omitempty"`

	// 用例包ID
	CasePackageId *string `json:"case_package_id,omitempty"`

	// 用例包名
	CasePackageName *string `json:"case_package_name,omitempty"`

	// 用例总数
	CaseTotal *int32 `json:"case_total,omitempty"`

	// 客户端类型，deviceTest使用
	ClientType *string `json:"client_type,omitempty"`

	CloudTestSuiteBasicInfo *CloudTestSuiteBasicInfo `json:"cloudTestSuite_basicInfo,omitempty"`

	Comments *string `json:"comments,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	// 环境类型：0表示用例包环境，1表示全局环境
	EnvType *int32 `json:"env_type,omitempty"`

	// environmentId环境信息
	EnvironmentGroupId *string `json:"environment_group_id,omitempty"`

	ExecuteStrategies *ExecuteStrategiesVo `json:"execute_strategies,omitempty"`

	// 用例类型
	ExecutorType *string `json:"executor_type,omitempty"`

	// 扩展参数
	ExtParams *[]TaskExtParam `json:"ext_params,omitempty"`

	Favorite *string `json:"favorite,omitempty"`

	// 唯一ID，主键
	Id *string `json:"id,omitempty"`

	IpGroup *IpGroup `json:"ipGroup,omitempty"`

	// 小网拨测替换application的hostIP
	IpKey *string `json:"ipKey,omitempty"`

	// 任务类型，是否debug任务
	IsDebugTask *int32 `json:"isDebugTask,omitempty"`

	// 执行标签
	Label *string `json:"label,omitempty"`

	// 资源池名称
	LabelName *string `json:"labelName,omitempty"`

	// 资源池类型
	LabelType *string `json:"labelType,omitempty"`

	// 最近一次停止时间
	LastStopTime *int64 `json:"lastStopTime,omitempty"`

	// 执行区域，冗余处理，修改更新在执行配置字段
	LocationIds *[]string `json:"location_ids,omitempty"`

	// 任务名
	Name *string `json:"name,omitempty"`

	// 测试计划Id
	PlanId *string `json:"planId,omitempty"`

	PreTestCaseInfo *PreTestCaseInfo `json:"preTestCaseInfo,omitempty"`

	ResourcePool *ResourcePool `json:"resourcePool,omitempty"`

	// 任务状态
	State *int32 `json:"state,omitempty"`

	SubtaskId *string `json:"subtask_id,omitempty"`

	// 子任务总数
	SubtaskTotal *int64 `json:"subtaskTotal,omitempty"`

	// 任务类型
	TaskTypeId *string `json:"taskTypeId,omitempty"`

	// 测试用例列表
	TestCases *[]TestCaseBasicInfo `json:"testCases,omitempty"`

	// 项目id
	TestServiceId *string `json:"test_service_id,omitempty"`

	// 测试套类型
	TestSuiteType *int32 `json:"testSuiteType,omitempty"`

	// 提示信息，用于任务操作过程中需要提供给前端的提示信息
	Tip *string `json:"tip,omitempty"`

	// 创建时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 更新人
	UpdateUser *string `json:"update_user,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`
}

func (o TaskInfoV4VoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInfoV4VoReq struct{}"
	}

	return strings.Join([]string{"TaskInfoV4VoReq", string(data)}, " ")
}
