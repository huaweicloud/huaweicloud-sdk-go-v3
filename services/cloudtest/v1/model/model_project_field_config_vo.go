package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectFieldConfigVo 实际的数据类型：单个对象，集合 或 NULL
type ProjectFieldConfigVo struct {

	// 字段配置URI标识
	Uri *string `json:"uri,omitempty"`

	Updator *NameAndIdVo `json:"updator,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 项目用例自定义字段名称
	CustomName *string `json:"customName,omitempty"`

	// 字段名（对应后端参数名）
	TableFieldName *string `json:"table_field_name,omitempty"`

	// 字段类型(单行文本text、多行文本textArea、单选框radio、多选框checkbox、日期date、数字number、用户user)。
	ValueType *string `json:"value_type,omitempty"`

	// 字段类型国际化名称
	ValueTypeName *string `json:"value_type_name,omitempty"`

	// 是否系统字段
	IsSystem *int32 `json:"is_system,omitempty"`

	// 是否显示
	IsDisplay *int32 `json:"is_display,omitempty"`

	// 是否必填
	IsRequired *int32 `json:"is_required,omitempty"`

	// 顺序数值
	SortNumb *int32 `json:"sort_numb,omitempty"`

	// 默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 扩展字段uri(用于连表查扩展字段)
	CustomFieldUri *string `json:"custom_field_uri,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	Creator *NameAndIdVo `json:"creator,omitempty"`

	// 创建时间时间戳
	CreateTimeStamp *int64 `json:"create_time_stamp,omitempty"`

	// 更新时间时间戳
	UpdateTimeStamp *int64 `json:"update_time_stamp,omitempty"`

	// 项目ID
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 可选项
	OptionVos *[]ProjectFieldConfigOptionVo `json:"option_vos,omitempty"`

	// 项目用例自定义字段id（1-25数字）
	CustomFieldId *int32 `json:"custom_field_id,omitempty"`

	// 项目用例自定义字段名称
	CustomFieldName *string `json:"custom_field_name,omitempty"`

	// 项目用例自定义字段入参或者返回参数名称
	CustomFieldParam *string `json:"custom_field_param,omitempty"`
}

func (o ProjectFieldConfigVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectFieldConfigVo struct{}"
	}

	return strings.Join([]string{"ProjectFieldConfigVo", string(data)}, " ")
}
