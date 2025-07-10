package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmToIncidentRequestBody struct {

	// 告警id，多个以”，分隔“
	AlarmIds string `json:"alarm_ids"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 责任人（个人）
	Assignee *string `json:"assignee,omitempty"`

	// 责任人（排班角色）
	AssigneeRole *string `json:"assignee_role,omitempty"`

	// 责任人（排版场景）
	AssigneeScene *string `json:"assignee_scene,omitempty"`

	// 附件
	Attachment *string `json:"attachment,omitempty"`

	// 归属应用
	CurrentCloudServiceId string `json:"current_cloud_service_id"`

	// 事件描述
	Description string `json:"description"`

	// 是否变更事件
	IsChangeEvent *bool `json:"is_change_event,omitempty"`

	// 业务是否中断
	IsServiceInterrupt bool `json:"is_service_interrupt"`

	// 事件等级
	LevelId string `json:"level_id"`

	// region
	MtmRegion *string `json:"mtm_region,omitempty"`

	// 事件类别
	MtmType string `json:"mtm_type"`

	// 资源id
	SourceId *string `json:"source_id,omitempty"`

	// 事件名称
	Title string `json:"title"`
}

func (o AlarmToIncidentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmToIncidentRequestBody struct{}"
	}

	return strings.Join([]string{"AlarmToIncidentRequestBody", string(data)}, " ")
}
