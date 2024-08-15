package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantWarRoomRequestBody 查询WarRoom信息请求
type ListTenantWarRoomRequestBody struct {

	// limit
	Limit *int64 `json:"limit,omitempty"`

	// 查询数量 最小值0 最大值1000
	Offset *int64 `json:"offset,omitempty"`

	// 用户id
	CurrentUsers *[]string `json:"current_users,omitempty"`

	// WarRoom单号，当有这个筛选条件时，其他筛选条件忽略
	WarRoomNums *[]string `json:"war_room_nums,omitempty"`

	// 事件单号 精确查询
	IncidentNum *string `json:"incident_num,omitempty"`

	// WarRoom名称 模糊查询
	Title *string `json:"title,omitempty"`

	// 区域 多选
	RegionCodeList *[]string `json:"region_code_list,omitempty"`

	// 事件级别 多选
	IncidentLevels *[]string `json:"incident_levels,omitempty"`

	// 影响应用id
	ImpactedApplicationIds *[]string `json:"impacted_application_ids,omitempty"`

	// WarRoom管理员
	Admin *[]string `json:"admin,omitempty"`

	// WarRoom状态
	Status *[]string `json:"status,omitempty"`

	// 拉起开始时间 默认前30天
	TriggeredStartTime *int64 `json:"triggered_start_time,omitempty"`

	// 拉起结束时间 默认当前时间
	TriggeredEndTime *int64 `json:"triggered_end_time,omitempty"`

	// 发生开始时间
	OccurStartTime *int64 `json:"occur_start_time,omitempty"`

	// 发生结束时间
	OccurEndTime *int64 `json:"occur_end_time,omitempty"`

	// 恢复开始时间
	RecoverStartTime *int64 `json:"recover_start_time,omitempty"`

	// 恢复结束时间
	RecoverEndTime *int64 `json:"recover_end_time,omitempty"`

	// 通报级别
	NotificationLevel *[]string `json:"notification_level,omitempty"`

	// 企业项目id
	EnterpriseProjectIds *[]string `json:"enterprise_project_ids,omitempty"`

	// WarRoom 单号 前端使用
	WarRoomNum *string `json:"war_room_num,omitempty"`

	// 是否统计,false 返回基本信息;true接口只返回统计结果：total_num,running_num,closed_num
	StatisticFlag *bool `json:"statistic_flag,omitempty"`
}

func (o ListTenantWarRoomRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantWarRoomRequestBody struct{}"
	}

	return strings.Join([]string{"ListTenantWarRoomRequestBody", string(data)}, " ")
}
