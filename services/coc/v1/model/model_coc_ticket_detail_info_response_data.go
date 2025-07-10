package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocTicketDetailInfoResponseData CocTicketDetailInfoResponseData。
type CocTicketDetailInfoResponseData struct {

	// 问题关联SLA。
	IssueCorrelationSla *string `json:"issue_correlation_sla,omitempty"`

	// 问题单级别。
	Level *string `json:"level,omitempty"`

	// 问题单责任服务。
	RootCauseCloudService *string `json:"root_cause_cloud_service,omitempty"`

	// 问题单根因分类。
	RootCauseType *string `json:"root_cause_type,omitempty"`

	// 问题单服务。
	CurrentCloudServiceId *string `json:"current_cloud_service_id,omitempty"`

	// 问题单接口人。
	IssueContactPerson *string `json:"issue_contact_person,omitempty"`

	// 发现问题版本号。
	IssueVersion *string `json:"issue_version,omitempty"`

	// 问题单来源。
	Source *string `json:"source,omitempty"`

	// 上传附件。
	CommitUploadAttachment *string `json:"commit_upload_attachment,omitempty"`

	// 问题单来源id。
	SourceId *string `json:"source_id,omitempty"`

	// 问题单企业项目。
	EnterpriseProject *string `json:"enterprise_project,omitempty"`

	// 问题单排班类型。
	VirtualScheduleType *string `json:"virtual_schedule_type,omitempty"`

	// 问题单标题。
	Title *string `json:"title,omitempty"`

	// 问题单region信息。
	Regions *string `json:"regions,omitempty"`

	// 问题单描述。
	Description *string `json:"description,omitempty"`

	// 问题单根因分析。
	RootCauseComment *string `json:"root_cause_comment,omitempty"`

	// 问题单解决方案。
	Solution *string `json:"solution,omitempty"`

	// 问题单区域搜索。
	RegionsSerch *string `json:"regions_serch,omitempty"`

	// 问题单级别审批配置。
	LevelApproveConfig *string `json:"level_approve_config,omitempty"`

	// 问题单挂起审批配置。
	SuspensionApproveConfig *string `json:"suspension_approve_config,omitempty"`

	// 处理时长。
	HandleTime *int64 `json:"handle_time,omitempty"`

	// 发现时间。
	FountTime *int64 `json:"fount_time,omitempty"`

	// 是否共性问题。
	IsCommonIssue *bool `json:"is_common_issue,omitempty"`

	// 问题单是否需要变更。
	IsNeedChange *bool `json:"is_need_change,omitempty"`

	// 问题单是否开启挂起。
	IsEnableSuspension *bool `json:"is_enable_suspension,omitempty"`

	// 是否启动异步流程。
	IsStartProcessAsync *bool `json:"is_start_process_async,omitempty"`

	// 是否重新提交空字段。
	IsUpdateNull *bool `json:"is_update_null,omitempty"`

	// 问题单创建人。
	Creator *string `json:"creator,omitempty"`

	// 问题单操作人。
	Operator *string `json:"operator,omitempty"`

	// 是否返回所有字段信息。
	IsReturnFullInfo *bool `json:"is_return_full_info,omitempty"`

	// 是否启动流程。
	IsStartProcess *bool `json:"is_start_process,omitempty"`

	// 问题单工单ID。
	TicketId *string `json:"ticket_id,omitempty"`

	// 问题单工单单号。
	RealTicketId *string `json:"real_ticket_id,omitempty"`

	// 问题单当前责任人。
	Assignee *string `json:"assignee,omitempty"`

	// 问题单参与人。
	Participator *string `json:"participator,omitempty"`

	// 问题单状态。
	WorkFlowStatus *string `json:"work_flow_status,omitempty"`

	// 流程状态。
	EngineErrorMsg *string `json:"engine_error_msg,omitempty"`

	// 基线状态。
	BaselineStatus *string `json:"baseline_status,omitempty"`

	// 工单类型。
	TicketType *string `json:"ticket_type,omitempty"`

	// 问题单当前阶段信息。
	Phase *string `json:"phase,omitempty"`

	// 变更子单信息。
	SubTickets *[]CocTicketDetailInfoResponseDataSubTickets `json:"sub_tickets,omitempty"`

	// 枚举列表。
	EnumDataList *[]IssuesTicketInfoEnumData `json:"enum_data_list,omitempty"`

	// 主键uuid
	Id *string `json:"id,omitempty"`

	// 应用版本
	MetaDataVersion *int32 `json:"meta_data_version,omitempty"`

	// 更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 创单时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 工单是否被删除。
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// 工单类型。
	TicketTypeId *string `json:"ticket_type_id,omitempty"`

	// 动作信息。
	FormInfo *interface{} `json:"_form_info,omitempty"`
}

func (o CocTicketDetailInfoResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocTicketDetailInfoResponseData struct{}"
	}

	return strings.Join([]string{"CocTicketDetailInfoResponseData", string(data)}, " ")
}
