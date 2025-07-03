package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocTicketHistoryBaseInfoV2 单据历史信息基础信息
type CocTicketHistoryBaseInfoV2 struct {

	// action
	Action *string `json:"action,omitempty"`

	// action中文名称
	ActionNameZh *string `json:"action_name_zh,omitempty"`

	// action英文名称
	ActionNameEn *string `json:"action_name_en,omitempty"`

	// 操作人ID
	Operator *string `json:"operator,omitempty"`

	// 当前状态
	Status *string `json:"status,omitempty"`

	// 操作开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 操作结束时间
	StopTime *string `json:"stop_time,omitempty"`

	// 备注信息
	Comment *string `json:"comment,omitempty"`

	// 枚举数据列表
	EnumDataList *[]CocTicketEnumDataV2 `json:"enum_data_list,omitempty"`
}

func (o CocTicketHistoryBaseInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocTicketHistoryBaseInfoV2 struct{}"
	}

	return strings.Join([]string{"CocTicketHistoryBaseInfoV2", string(data)}, " ")
}
