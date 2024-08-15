package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnumDataInfo 基础数据
type EnumDataInfo struct {

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

	// 是否删除
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// 匹配类型
	MatchType *string `json:"match_type,omitempty"`

	// 所属单号
	TicketId *string `json:"ticket_id,omitempty"`

	// 中文名
	NameZh *string `json:"name_zh,omitempty"`

	// 英文名
	NameEn *string `json:"name_en,omitempty"`

	// 工号
	UserName *string `json:"user_name,omitempty"`

	// 中文名
	FullNameZh *string `json:"full_name_zh,omitempty"`

	// 英文名
	FullNameEn *string `json:"full_name_en,omitempty"`

	// 中文名路径
	NameZhPath *string `json:"name_zh_path,omitempty"`

	// 中文名路径
	NameEnPath *string `json:"name_en_path,omitempty"`

	// status
	Status *string `json:"status,omitempty"`

	// 唯一id
	BizId *string `json:"biz_id,omitempty"`

	// 字段id
	PropId *string `json:"prop_id,omitempty"`

	// 模型id
	ModelId *string `json:"model_id,omitempty"`

	// 枚举类型
	EnumTypeId *string `json:"enum_type_id,omitempty"`
}

func (o EnumDataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnumDataInfo struct{}"
	}

	return strings.Join([]string{"EnumDataInfo", string(data)}, " ")
}
