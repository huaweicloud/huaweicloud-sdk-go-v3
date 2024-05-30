package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StandElementFieldVo 属性。
type StandElementFieldVo struct {

	// 属性名称。
	FdName string `json:"fd_name"`

	// 属性英文名称。
	FdNameEn *string `json:"fd_name_en,omitempty"`

	// 属性描述。
	Description *string `json:"description,omitempty"`

	// 属性英文描述。
	DescriptionEn *string `json:"descriptionEn,omitempty"`

	// 属性标签。
	Label *string `json:"label,omitempty"`

	// 是否禁用。
	Disabled *bool `json:"disabled,omitempty"`

	// 数据标准ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 是否显示，系统默认项必然显示不允许修改。true表示使用数据标准时体现（增改查的时候可以操作该属性），false表示使用数据标准时不体现。
	Actived bool `json:"actived"`

	// 是否必填。true：必填，false：非必填。
	Required *bool `json:"required,omitempty"`

	// 是否可搜索。true表示在数据标准列表页面可搜索，false表示在数据标准列表页面不可搜索。
	Searchable *bool `json:"searchable,omitempty"`

	// 允许值。
	OptionalValues *string `json:"optional_values,omitempty"`

	// 字段类型，0表示系统字段， 1表示自定义字段。
	FieldType *int32 `json:"field_type,omitempty"`

	// 前端展示名。
	DisplayedName *string `json:"displayed_name,omitempty"`

	// 前端展示名英文。
	DisplayedNameEn *string `json:"displayed_name_en,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`
}

func (o StandElementFieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandElementFieldVo struct{}"
	}

	return strings.Join([]string{"StandElementFieldVo", string(data)}, " ")
}
