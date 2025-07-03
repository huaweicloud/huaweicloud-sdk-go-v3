package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocIncidentDataV2 IncidentData
type CocIncidentDataV2 struct {

	// 云服务
	CurrentCloudServiceId *string `json:"current_cloud_service_id,omitempty"`

	// 事件等级
	LevelId *string `json:"level_id,omitempty"`

	// 区域Region
	MtmRegion *string `json:"mtm_region,omitempty"`

	// 事件来源
	SourceId *string `json:"source_id,omitempty"`

	// 转发规则
	ForwardRuleId *string `json:"forward_rule_id,omitempty"`

	// 企业应用
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 事件类别
	MtmType *string `json:"mtm_type,omitempty"`

	// 事件标题
	Title *string `json:"title,omitempty"`

	// 事件描述
	Description *string `json:"description,omitempty"`

	// 事件单号
	TicketId *string `json:"ticket_id,omitempty"`

	// 服务是否中断
	IsServiceInterrupt *string `json:"is_service_interrupt,omitempty"`

	// 流程状态
	WorkFlowStatus *string `json:"work_flow_status,omitempty"`

	// 阶段
	Phase *string `json:"phase,omitempty"`

	// 责任人
	Assignee *string `json:"assignee,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 创建人
	Operator *string `json:"operator,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 故障开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 最近一次处理时间
	HandleTime *string `json:"handle_time,omitempty"`

	// 事件归属
	IncidentOwnership *string `json:"incident_ownership,omitempty"`

	// 枚举列表
	EnumDataList *[]CocTicketInfoEnumDataV2 `json:"enum_data_list,omitempty"`
}

func (o CocIncidentDataV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocIncidentDataV2 struct{}"
	}

	return strings.Join([]string{"CocIncidentDataV2", string(data)}, " ")
}
