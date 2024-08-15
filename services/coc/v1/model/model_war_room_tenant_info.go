package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WarRoomTenantInfo WarRoom信息
type WarRoomTenantInfo struct {

	// 主键
	Id *string `json:"id,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// WarRoom管理员
	Admin *string `json:"admin,omitempty"`

	// 恢复成员
	RecoverMember *[]string `json:"recover_member,omitempty"`

	// 主恢复责任人
	RecoverLeader *[]string `json:"recover_leader,omitempty"`

	Incident *WarRoomIncident `json:"incident,omitempty"`

	// 事件来源
	Source *string `json:"source,omitempty"`

	// 影响的Region
	Regions *[]WarRoomTenantInfoRegions `json:"regions,omitempty"`

	// 变更单号
	ChangeNum *string `json:"change_num,omitempty"`

	// 开始时间
	OccurTime *int64 `json:"occur_time,omitempty"`

	// 故障恢复时间
	RecoverTime *int64 `json:"recover_time,omitempty"`

	// 故障原因
	FaultCause *string `json:"fault_cause,omitempty"`

	// 添加时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 首次通报时间
	FirstReportTime *int64 `json:"first_report_time,omitempty"`

	// 恢复通报时间
	RecoveryNotificationTime *int64 `json:"recovery_notification_time,omitempty"`

	// 故障影响
	FaultImpact *string `json:"fault_impact,omitempty"`

	// WarRoom描述
	Description *string `json:"description,omitempty"`

	// 通报级别 租户区同事件级别
	CircularLevel *string `json:"circular_level,omitempty"`

	WarRoomStatus *WarRoomEnumeration `json:"war_room_status,omitempty"`

	// 影响应用
	ImpactedApplication *[]WarRoomTenantInfoImpactedApplication `json:"impacted_application,omitempty"`

	// 处理时长(分钟)
	ProcessingDuration *int64 `json:"processing_duration,omitempty"`

	// 恢复时长(分钟)
	RestorationDuration *int64 `json:"restoration_duration,omitempty"`

	// WarRoom单号
	WarRoomNum *string `json:"war_room_num,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o WarRoomTenantInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WarRoomTenantInfo struct{}"
	}

	return strings.Join([]string{"WarRoomTenantInfo", string(data)}, " ")
}
