package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TicketHistoryInfo struct {

	// 扩展字段
	CurrentCloudServiceId *string `json:"current_cloud_service_id,omitempty"`

	// 扩展字段
	Description *string `json:"description,omitempty"`

	// 扩展字段
	LevelId *string `json:"level_id,omitempty"`

	// 扩展字段
	MtmRegion *string `json:"mtm_region,omitempty"`

	// 扩展字段
	MtmType *string `json:"mtm_type,omitempty"`

	// 扩展字段
	SourceId *string `json:"source_id,omitempty"`

	// 扩展字段
	Title *string `json:"title,omitempty"`

	// 是否变更事件
	IsChangeEvent *bool `json:"is_change_event,omitempty"`

	// 是否变更事件
	IsServiceInterrupt *bool `json:"is_service_interrupt,omitempty"`

	// 操作标识
	ActionId *string `json:"action_id,omitempty"`

	// 动作
	Action *string `json:"action,omitempty"`

	// 子动作
	SubAction *string `json:"sub_action,omitempty"`

	// 操作人
	Operator *string `json:"operator,omitempty"`

	// 评论
	Comment *string `json:"comment,omitempty"`

	// 主键
	Id *string `json:"id,omitempty"`

	// 单号
	TicketId *string `json:"ticket_id,omitempty"`

	// 起始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	StopTime *int64 `json:"stop_time,omitempty"`

	// 对象类型
	TargetType *string `json:"target_type,omitempty"`

	// 对象值
	TargetValue *string `json:"target_value,omitempty"`

	// 待修改
	IsDeteted *bool `json:"is_deteted,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// action中文名
	ActionNameZh *string `json:"action_name_zh,omitempty"`

	// action英文名
	ActionNameEn *string `json:"action_name_en,omitempty"`

	// action中文模板
	ActionTemplateZh *string `json:"action_template_zh,omitempty"`

	// action中文模板
	ActionTemplateEn *string `json:"action_template_en,omitempty"`

	// 工单状态
	Status *string `json:"status,omitempty"`

	// 最终子动作
	FinalSubAction *string `json:"final_sub_action,omitempty"`

	// 枚举数据
	EnumDataList *[]EnumDataInfo `json:"enum_data_list,omitempty"`
}

func (o TicketHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TicketHistoryInfo struct{}"
	}

	return strings.Join([]string{"TicketHistoryInfo", string(data)}, " ")
}
