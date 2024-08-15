package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IncidentTicketInfoResponseData IncidentTicketInfoResponseData
type IncidentTicketInfoResponseData struct {

	// 事件单号
	IncidentNum *string `json:"incident_num,omitempty"`

	// 区域Code，如果自动拉起WarRoom则为必填，现在只支持1个
	Region *[]string `json:"region,omitempty"`

	// 企业项目ID，现在只支持1个
	EnterpriseProject *[]string `json:"enterprise_project,omitempty"`

	// 归属应用ID，现在只支持1个
	CurrentCloudService *[]string `json:"current_cloud_service,omitempty"`

	// 事件级别 参考：枚举 事件级别incident_level
	IncidentLevel *string `json:"incident_level,omitempty"`

	// 业务是否中断，取值：true/false
	IsServiceInterrupt *bool `json:"is_service_interrupt,omitempty"`

	// 事件类别 参考：枚举 事件类别incident_type
	IncidentType *string `json:"incident_type,omitempty"`

	// 事件标题，最大长度：200
	IncidentTitle *string `json:"incident_title,omitempty"`

	// 事件描述，最大长度：600
	IncidentDescription *string `json:"incident_description,omitempty"`

	// 单据来源 参考：枚举 事件来源incident_source
	IncidentSource *string `json:"incident_source,omitempty"`

	// 责任人，排班场景和排班角色不能同时为空，现在只支持1个
	IncidentAssignee *[]string `json:"incident_assignee,omitempty"`

	// 排班场景，责任人和排班角色不能同时为空
	AssigneeScene *string `json:"assignee_scene,omitempty"`

	// 排班角色，排班场景和责任人不能同时为空
	AssigneeRole *string `json:"assignee_role,omitempty"`

	// warroom_id
	WarroomId *string `json:"warroom_id,omitempty"`

	// 最后一次提交解决方案时间戳
	HandleTime *int64 `json:"handle_time,omitempty"`

	// 状态KEY
	Status *string `json:"status,omitempty"`

	// 创单时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 创单人
	Creator *string `json:"creator,omitempty"`

	// 枚举列表
	EnumDataList *[]TicketInfoEnumData `json:"enum_data_list,omitempty"`
}

func (o IncidentTicketInfoResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentTicketInfoResponseData struct{}"
	}

	return strings.Join([]string{"IncidentTicketInfoResponseData", string(data)}, " ")
}
