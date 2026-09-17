package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskInfo 测试套件信息
type TaskInfo struct {

	// 指定创建任务的uri
	Uri *string `json:"uri,omitempty"`

	// 分支/迭代uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 处理人/责任人id
	OwnerId *string `json:"owner_id,omitempty"`

	// 父任务uri
	ParentUri *string `json:"parent_uri,omitempty"`

	// 动态任务用例过滤条件
	TestCaseCondition *string `json:"test_case_condition,omitempty"`

	// 测试阶段
	Stage *string `json:"stage,omitempty"`

	// 服务类型0:功能测试 1:接口测试 11:性能测试
	ServiceType *int32 `json:"service_type,omitempty"`

	// 编号
	Number *string `json:"number,omitempty"`

	// 标记id
	Tags *[]string `json:"tags,omitempty"`

	// 模块id
	ModuleId *string `json:"module_id,omitempty"`

	// 模块名称
	ModuleName *string `json:"module_name,omitempty"`

	// 发布版本号
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 状态code
	StatusCode *int32 `json:"status_code,omitempty"`

	// 扩展参数
	ExtParam *string `json:"ext_param,omitempty"`

	// 执行方式 1：串行，2：并行
	ExecuteWay *int32 `json:"execute_way,omitempty"`

	// 执行类型（0：冒烟，1：定时）
	ExecuteType *int32 `json:"execute_type,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 计划开始时间戳，当传入-1时，时间置为空
	PlanStartTimestamp *int64 `json:"plan_start_timestamp,omitempty"`

	// 计划结束时间戳，当传入-1时，时间置为空
	PlanEndTimestamp *int64 `json:"plan_end_timestamp,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 任务关联用例uri数组，CloudDragon环境
	AssignCaseUris *[]string `json:"assign_case_uris,omitempty"`

	CaseOperationInfo *CaseOperationInfo `json:"case_operation_info,omitempty"`

	// 是否只需要修改测试套状态
	OnlyUpdateStatus *bool `json:"only_update_status,omitempty"`

	// 是否异步
	IsAsync *bool `json:"is_async,omitempty"`
}

func (o TaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInfo struct{}"
	}

	return strings.Join([]string{"TaskInfo", string(data)}, " ")
}
