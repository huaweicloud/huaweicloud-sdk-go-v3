package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IteratorVersionSummaryVo 实际的数据类型：单个对象，集合 或 NULL
type IteratorVersionSummaryVo struct {

	// 资源URI
	Uri *string `json:"uri,omitempty"`

	// 资源类型
	Type *string `json:"type,omitempty"`

	// 创建人
	Author *string `json:"author,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 级别
	Rank *int32 `json:"rank,omitempty"`

	// 待测版本
	Version *string `json:"version,omitempty"`

	// 处理者ID
	Owner *string `json:"owner,omitempty"`

	// 创建人ID
	Creator *string `json:"creator,omitempty"`

	// 关联迭代
	Iterations *string `json:"iterations,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 最后修改人
	LastModifier *string `json:"last_modifier,omitempty"`

	// 最后修改时间
	LastModified *sdktime.SdkTime `json:"last_modified,omitempty"`

	// 修改时间时间戳
	LastModifiedTimestamp *int64 `json:"last_modified_timestamp,omitempty"`

	// 最后变更时间
	LastChangeTime *sdktime.SdkTime `json:"last_change_time,omitempty"`

	// 版本URI
	VersionUri *string `json:"version_uri,omitempty"`

	// 源资源URI
	OriginUri *string `json:"origin_uri,omitempty"`

	// 父资源URI
	ParentUri *string `json:"parent_uri,omitempty"`

	// 父资源路径
	ParentPath *string `json:"parent_path,omitempty"`

	// 创建版本URI
	CreationVersionUri *string `json:"creation_version_uri,omitempty"`

	// 创建时间
	CreationDate *sdktime.SdkTime `json:"creation_date,omitempty"`

	// 创建时间时间戳
	CreationDateTimestamp *int64 `json:"creation_date_timestamp,omitempty"`

	// 创建人名称
	AuthorName *string `json:"author_name,omitempty"`

	// 备注
	Comment *string `json:"comment,omitempty"`

	// 编号
	Number *string `json:"number,omitempty"`

	// 是否为Master分支
	IsMaster *int32 `json:"is_master,omitempty"`

	// 是否为迭代
	IsIterator *int32 `json:"is_iterator,omitempty"`

	// 开始时间
	PlanStartDate *sdktime.SdkTime `json:"plan_start_date,omitempty"`

	// 结束时间
	PlanEndDate *sdktime.SdkTime `json:"plan_end_date,omitempty"`

	// 微服务ID
	ServiceId *string `json:"service_id,omitempty"`

	// 微服务名
	ServiceName *string `json:"service_name,omitempty"`

	// PBI ID
	PbiId *string `json:"pbi_id,omitempty"`

	// PBI信息
	PbiName *string `json:"pbi_name,omitempty"`

	// 计划ID
	PlanId *string `json:"plan_id,omitempty"`

	// 度量PBI ID
	MetricPbiIds *string `json:"metric_pbi_ids,omitempty"`

	// 度量PBI名称
	MetricPbiIdNames *string `json:"metric_pbi_id_names,omitempty"`

	// 最后同步时间
	LastSynDate *sdktime.SdkTime `json:"last_syn_date,omitempty"`

	// 版本是否关闭
	IsClosed *string `json:"is_closed,omitempty"`

	// 是否同步git库
	AsynGit *string `json:"asyn_git,omitempty"`

	// schema编号
	SchemaNo *int32 `json:"schema_no,omitempty"`

	// 迭代实际完成时间
	FinishDate *sdktime.SdkTime `json:"finish_date,omitempty"`

	// 处理者名称
	OwnerName *string `json:"owner_name,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 当前所处阶段
	CurrentStage *string `json:"current_stage,omitempty"`

	// 服务类型
	ServiceTypes *string `json:"service_types,omitempty"`

	// 风险等级
	RiskRating *int32 `json:"risk_rating,omitempty"`

	// 风险描述
	RiskDes *string `json:"risk_des,omitempty"`

	// 项目ID
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// pi的id
	PiId *string `json:"pi_id,omitempty"`

	// 计划开始时间
	StartDate *string `json:"start_date,omitempty"`

	// 计划开始时间时间戳
	StartDateTimestamp *int64 `json:"start_date_timestamp,omitempty"`

	// 计划结束时间
	EndDate *string `json:"end_date,omitempty"`

	// 计划结束时间时间戳
	EndDateTimestamp *int64 `json:"end_date_timestamp,omitempty"`

	// 实际开始时间
	ActualStartDate *string `json:"actual_start_date,omitempty"`

	// 实际开始时间时间戳
	ActualStartDateTimestamp *int64 `json:"actual_start_date_timestamp,omitempty"`

	// 实际完成时间
	ActualEndDate *string `json:"actual_end_date,omitempty"`

	// 实际开始时间时间戳
	ActualEndDateTimestamp *int64 `json:"actual_end_date_timestamp,omitempty"`

	// 是否超期
	IsExpired *string `json:"is_expired,omitempty"`

	// 计划过期信息,空代表不超期，否则给出具体超期信息
	ExpiredInfo *string `json:"expired_info,omitempty"`

	// 迭代计划，默认包含design,execute,report
	Stages *[]string `json:"stages,omitempty"`

	Design *DesignSummaryVo `json:"design,omitempty"`

	Execute *ExecuteSummaryVo `json:"execute,omitempty"`

	Report *ReportSummaryVo `json:"report,omitempty"`

	// 所属分支URI
	BranchUri *string `json:"branch_uri,omitempty"`

	// 所属分支名称
	BranchName *string `json:"branch_name,omitempty"`
}

func (o IteratorVersionSummaryVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IteratorVersionSummaryVo struct{}"
	}

	return strings.Join([]string{"IteratorVersionSummaryVo", string(data)}, " ")
}
