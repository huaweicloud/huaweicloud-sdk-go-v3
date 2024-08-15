package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateWarRoomRequestBody struct {

	// WarRoom名称
	WarRoomName string `json:"war_room_name"`

	// WarRoom描述
	Description *string `json:"description,omitempty"`

	// 区域id
	RegionCodeList *[]string `json:"region_code_list,omitempty"`

	// 影响应用id
	ApplicationIdList []string `json:"application_id_list"`

	// 事件单号
	IncidentNumber string `json:"incident_number"`

	// 排班分组
	ScheduleGroup []ScheduleGroupInfo `json:"schedule_group"`

	// 参与者
	Participant *[]string `json:"participant,omitempty"`

	// WarRoom管理员
	WarRoomAdmin string `json:"war_room_admin"`

	// 应用名称列表
	ApplicationNames *[]string `json:"application_names,omitempty"`

	// region名称列表
	RegionNames *[]string `json:"region_names,omitempty"`

	// 企业项目id
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 创建群组方式
	NotificationType *string `json:"notification_type,omitempty"`
}

func (o CreateWarRoomRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWarRoomRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWarRoomRequestBody", string(data)}, " ")
}
