package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskVo 测试任务集合
type TaskVo struct {

	// 测试任务URI
	Uri *string `json:"uri,omitempty"`

	// 测试任务名称
	Name *string `json:"name,omitempty"`

	// 测试阶段
	Stage *string `json:"stage,omitempty"`

	// 编号
	Number *string `json:"number,omitempty"`

	// 标签
	Tags *string `json:"tags,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 创建人ID
	AuthorId *string `json:"author_id,omitempty"`

	// 创建人名称
	AuthorName *string `json:"author_name,omitempty"`

	// 责任人ID
	OwnerId *string `json:"owner_id,omitempty"`

	// 责任人名称
	OwnerName *string `json:"owner_name,omitempty"`

	// 父任务URI
	ParentUri *string `json:"parent_uri,omitempty"`

	// 父任务路径
	ParentPath *string `json:"parent_path,omitempty"`

	// 源任务URI
	OriginUri *string `json:"origin_uri,omitempty"`

	// 版本URI
	VersionUri *string `json:"version_uri,omitempty"`

	// 分支URI
	BranchUri *string `json:"branch_uri,omitempty"`

	// 版本名称
	VersionName *string `json:"version_name,omitempty"`

	// 创建时间
	CreationDate *sdktime.SdkTime `json:"creation_date,omitempty"`

	// 创建时间时间戳
	CreateDateTimestamp *int64 `json:"create_date_timestamp,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 更新时间时间戳
	UpdateTimeTimestamp *int64 `json:"update_time_timestamp,omitempty"`

	// 关联关系修改时时间
	RelationChangeTime *sdktime.SdkTime `json:"relation_change_time,omitempty"`

	// 关联关系修改时间时间戳
	RelationChangeTimeTimestamp *int64 `json:"relation_change_time_timestamp,omitempty"`

	// 动态任务用例过滤条件
	TestCaseCondition *string `json:"test_case_condition,omitempty"`

	// 修改人Id
	UpdatorId *string `json:"updator_id,omitempty"`

	// 修改人名称
	UpdatorName *string `json:"updator_name,omitempty"`

	// 关联关系修改人Id
	RelationChangerId *string `json:"relation_changer_id,omitempty"`

	// 服务类型ID
	ServiceType *int32 `json:"service_type,omitempty"`

	// 服务类型名称
	ServiceTypeName *string `json:"service_type_name,omitempty"`

	// 标签名称集合
	TagList *[]string `json:"tag_list,omitempty"`

	// 模块ID
	ModuleId *string `json:"module_id,omitempty"`

	// 模块名称
	ModuleName *string `json:"module_name,omitempty"`

	// 模块path
	ModulePath *string `json:"module_path,omitempty"`

	// 模块路径名称
	ModulePathName *string `json:"module_path_name,omitempty"`

	// 发布版本号
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 扩展参数
	ExtParam *string `json:"ext_param,omitempty"`

	// 执行方式（1：串行，2：并行）
	ExecuteWay *int32 `json:"execute_way,omitempty"`

	// 执行类型（0：冒烟，1：定时）
	ExecuteType *int32 `json:"execute_type,omitempty"`

	// 生命周期状态Code
	StatusCode *int32 `json:"status_code,omitempty"`

	// 生命周期状态名称
	StatusName *string `json:"status_name,omitempty"`

	// 执行结果Code
	ResultCode *int32 `json:"result_code,omitempty"`

	// 执行状态名称
	ResultName *string `json:"result_name,omitempty"`

	// Echo执行状态Code
	ExecuteStatusCode *int32 `json:"execute_status_code,omitempty"`

	// Echo执行状态名称
	ExecuteStatusName *string `json:"execute_status_name,omitempty"`

	// 执行人ID
	ExecutorId *string `json:"executor_id,omitempty"`

	// 执行人名称
	ExecutorName *string `json:"executor_name,omitempty"`

	// 最近执行时间
	ExecuteLatestTime *sdktime.SdkTime `json:"execute_latest_time,omitempty"`

	// 最近执行时间时间戳
	ExecuteLatestTimeTimestamp *int64 `json:"execute_latest_time_timestamp,omitempty"`

	// 执行时长
	ExecuteDuration *string `json:"execute_duration,omitempty"`

	// 执行次数
	ExecuteTimes *int32 `json:"execute_times,omitempty"`

	// 项目ID
	ProjectUuid *string `json:"project_uuid,omitempty"`

	CaseOperationInfo *CaseOperationVo `json:"case_operation_info,omitempty"`

	// 关联用例数
	AssignCaseNum *int32 `json:"assign_case_num,omitempty"`

	// 已完成用例数量
	FinishCaseNum *int32 `json:"finish_case_num,omitempty"`

	// 关联缺陷数量
	AssignDefectNum *int32 `json:"assign_defect_num,omitempty"`

	// 任务关联用例变更提示信息
	TaskAssignMsg *string `json:"task_assign_msg,omitempty"`

	// 测试套所属迭代uri，非迭代下创建的测试套返回null
	IteratorVersionUri *string `json:"iterator_version_uri,omitempty"`

	// 用户自定义结果对应的用例数目
	ResultNumberList *[]NameAndValueAndCodeVo `json:"result_number_list,omitempty"`

	// 测试套完成时间
	FinishDate *sdktime.SdkTime `json:"finish_date,omitempty"`

	// 测试套完成时间戳
	FinishDateTimestamp *int64 `json:"finish_date_timestamp,omitempty"`

	// 计划开始时间
	PlanStartDate *sdktime.SdkTime `json:"plan_start_date,omitempty"`

	// 计划开始时间戳
	PlanStartTimestamp *int64 `json:"plan_start_timestamp,omitempty"`

	// 计划结束时间
	PlanEndDate *sdktime.SdkTime `json:"plan_end_date,omitempty"`

	// 计划结束时间戳
	PlanEndTimestamp *int64 `json:"plan_end_timestamp,omitempty"`

	// 测试套超期状态值，分别为：无状态(null)、未超期(0)、即将超期(1)、已超期(2)、延期完成(3)、按期完成(4)
	ExpirationStatus *int32 `json:"expiration_status,omitempty"`

	// 测试套超期状态名称，分别为：无状态(不显示状态)、未超期(Unexpired)、即将超期(About to expire)、已超期(Expired)、延期完成(Delayed completion)、按期完成(On schedule completion)
	ExpirationStatusName *string `json:"expiration_status_name,omitempty"`
}

func (o TaskVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskVo struct{}"
	}

	return strings.Join([]string{"TaskVo", string(data)}, " ")
}
