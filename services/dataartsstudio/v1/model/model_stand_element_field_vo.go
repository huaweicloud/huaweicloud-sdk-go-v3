package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StandElementFieldVo 属性
type StandElementFieldVo struct {

	// 属性名称
	FdName string `json:"fd_name"`

	// 属性描述
	Description *string `json:"description,omitempty"`

	// ID
	Id *int64 `json:"id,omitempty"`

	// 是否显示，系统默认项必然显示不允许修改
	Actived *bool `json:"actived,omitempty"`

	// 是否必填
	Required *bool `json:"required,omitempty"`

	// 是否可搜索
	Searchable *bool `json:"searchable,omitempty"`

	// 前端展示名
	DisplayedName *string `json:"displayed_name,omitempty"`

	// 前端展示名英文
	DisplayedNameEn *string `json:"displayed_name_en,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`
}

func (o StandElementFieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandElementFieldVo struct{}"
	}

	return strings.Join([]string{"StandElementFieldVo", string(data)}, " ")
}
