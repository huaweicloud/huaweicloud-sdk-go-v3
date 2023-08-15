package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIncident 事件详情
type CreateIncident struct {

	// 版本
	Version *string `json:"version,omitempty"`

	// 严重性等级
	Severity *string `json:"severity,omitempty"`

	Environment *ShowAlertRspEnvironment `json:"environment,omitempty"`

	DataSource *CreateAlertDataSource `json:"data_source,omitempty"`

	// Update time
	FirstObservedTime *string `json:"first_observed_time,omitempty"`

	// Update time
	LastObservedTime *string `json:"last_observed_time,omitempty"`

	// Create time
	CreateTime *string `json:"create_time,omitempty"`

	// Update time
	ArriveTime *string `json:"arrive_time,omitempty"`

	// The name, display only
	Title *string `json:"title,omitempty"`

	// The label, display only
	Labels *string `json:"labels,omitempty"`

	// The description, display only
	Description *string `json:"description,omitempty"`

	// 事件URL链接
	SourceUrl *string `json:"source_url,omitempty"`

	// 事件发生次数
	Count *int32 `json:"count,omitempty"`

	// 置信度
	Confidence *int32 `json:"confidence,omitempty"`

	// 严重性等级
	Serverity *string `json:"serverity,omitempty"`

	// 关键性，是指事件涉及的资源的重要性级别。
	Criticality *int32 `json:"criticality,omitempty"`

	IncidentType *CreateIncidentIncidentType `json:"incident_type,omitempty"`

	// network_list
	NetworkList *[]CreateIncidentNetworkList `json:"network_list,omitempty"`

	// network_list
	ResourceList *[]CreateIncidentResourceList `json:"resource_list,omitempty"`

	Remediation *ShowAlertRspRemediation `json:"remediation,omitempty"`

	// 验证状态
	VerificationState *string `json:"verification_state,omitempty"`

	// 事件处理状态
	HandleStatus *string `json:"handle_status,omitempty"`

	// sla
	Sla *int32 `json:"sla,omitempty"`

	// Create time
	UpdateTime *string `json:"update_time,omitempty"`

	// Create time
	CloseTime *string `json:"close_time,omitempty"`

	// 周期/处置阶段编号
	ChopPhase *string `json:"chop_phase,omitempty"`

	// 周期/处置阶段编号
	IpdrrPhase *string `json:"ipdrr_phase,omitempty"`

	// 周期/处置阶段编号
	PpdrPhase *string `json:"ppdr_phase,omitempty"`

	// 是否为调试事件.
	Simulation *string `json:"simulation,omitempty"`

	// 委托人
	Actor *string `json:"actor,omitempty"`

	// The name, display only
	Owner *string `json:"owner,omitempty"`

	// The name, display only
	Cteator *string `json:"cteator,omitempty"`

	// 关闭原因
	CloseReason *string `json:"close_reason,omitempty"`

	// 关闭原因
	CloseComment *string `json:"close_comment,omitempty"`

	Malware *CreateIncidentMalware `json:"malware,omitempty"`

	// 系统信息
	SystemInfo *interface{} `json:"system_info,omitempty"`

	// 进程信息
	Process *[]CreateIncidentProcess `json:"process,omitempty"`

	// 用户信息
	UserInfo *[]CreateIncidentUserInfo `json:"user_info,omitempty"`

	// 文件信息
	FileInfo *[]ShowAlertRspFileInfo `json:"file_info,omitempty"`

	// 系统信息
	SystemIncidentTable *interface{} `json:"system_incident_table,omitempty"`

	// Id value
	Id *string `json:"id,omitempty"`

	// workspace id
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o CreateIncident) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIncident struct{}"
	}

	return strings.Join([]string{"CreateIncident", string(data)}, " ")
}
