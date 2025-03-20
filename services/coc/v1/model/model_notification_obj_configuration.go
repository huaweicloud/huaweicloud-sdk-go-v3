package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationObjConfiguration 通知对象配置
type NotificationObjConfiguration struct {

	// 通知对象类型（排班/个人/工单责任人/群组）
	NotificationEndpointType *string `json:"notification_endpoint_type,omitempty"`

	// 排班场景ID
	ScheduleScene *string `json:"schedule_scene,omitempty"`

	// 排班角色ID
	ScheduleRole *string `json:"schedule_role,omitempty"`

	// 排班角色名称
	ScheduleRoleName *string `json:"schedule_role_name,omitempty"`

	// 消息通知接收人，对于群组通知而言其对应的是被@的群成员
	Recipients *string `json:"recipients,omitempty"`

	// 群组类型（企业微信/钉钉/飞书）
	GroupType *string `json:"group_type,omitempty"`

	// 群组ID
	GroupId *string `json:"group_id,omitempty"`

	// 群组名称
	GroupName *string `json:"group_name,omitempty"`
}

func (o NotificationObjConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationObjConfiguration struct{}"
	}

	return strings.Join([]string{"NotificationObjConfiguration", string(data)}, " ")
}
