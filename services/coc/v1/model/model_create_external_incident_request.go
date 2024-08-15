package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExternalIncidentRequest CreateExternalIncidentRequest 创单请求参数
type CreateExternalIncidentRequest struct {

	// 区域Code，如果自动拉起WarRoom则为必填，现在只支持1个
	Region *[]string `json:"region,omitempty"`

	// 企业项目ID，现在只支持1个
	EnterpriseProject *[]string `json:"enterprise_project,omitempty"`

	// 归属应用ID，现在只支持1个
	CurrentCloudService *[]string `json:"current_cloud_service,omitempty"`

	// 事件级别 参考：枚举 事件级别incident_level
	IncidentLevel string `json:"incident_level"`

	// 业务是否中断，取值：true/false
	IsServiceInterrupt bool `json:"is_service_interrupt"`

	// 事件类别 参考：枚举 事件类别incident_type
	IncidentType string `json:"incident_type"`

	// 事件标题，最大长度：200
	IncidentTitle string `json:"incident_title"`

	// 事件描述，最大长度：600
	IncidentDescription *string `json:"incident_description,omitempty"`

	// 单据来源 参考：枚举 事件来源incident_source
	IncidentSource string `json:"incident_source"`

	// 责任人，排班场景和排班角色不能同时为空，现在只支持1个
	IncidentAssignee *[]string `json:"incident_assignee,omitempty"`

	// 排班场景，责任人和排班角色不能同时为空
	AssigneeScene *string `json:"assignee_scene,omitempty"`

	// 排班角色，排班场景和责任人不能同时为空
	AssigneeRole *string `json:"assignee_role,omitempty"`

	// 创单人
	Creator string `json:"creator"`
}

func (o CreateExternalIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalIncidentRequest struct{}"
	}

	return strings.Join([]string{"CreateExternalIncidentRequest", string(data)}, " ")
}
