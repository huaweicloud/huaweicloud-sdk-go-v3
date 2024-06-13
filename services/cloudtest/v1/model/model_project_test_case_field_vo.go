package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectTestCaseFieldVo struct {

	// 项目用例自定义字段主键
	Uri *string `json:"uri,omitempty"`

	// 项目用例自定义字段名称
	Name *string `json:"name,omitempty"`

	// 项目用例自定义字段类型（单行文本text、多行文本textArea、单选框radio、多选框checkBox、日期date、数字number、单选用户user）
	Type *string `json:"type,omitempty"`

	// 项目用例自定义字段选项（数字类型时，数组两个值，第一个是最小值，第二个是最大值）
	Options *string `json:"options,omitempty"`

	// 项目用例自定义字段描述
	Description *string `json:"description,omitempty"`

	// 项目用例自定义字段创建人
	Creator *string `json:"creator,omitempty"`

	// 项目用例自定义字段更新人
	Updater *string `json:"updater,omitempty"`

	// 项目用例自定义字段id（1-25数字）
	CustomFieldId *int32 `json:"custom_field_id,omitempty"`

	// 项目用例自定义字段名称
	CustomFieldName *string `json:"custom_field_name,omitempty"`

	// 项目用例自定义字段入参或者返回参数名称
	CustomFieldParam *string `json:"custom_field_param,omitempty"`

	// 项目用例自定义字段类型国际化名称
	TypeName *string `json:"type_name,omitempty"`

	// 项目用例自定义字段创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 项目用例自定义字段创建时间时间戳
	CreateTimeTimestamp *int64 `json:"create_time_timestamp,omitempty"`

	// 项目用例自定义字段更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 项目用例自定义字段更新时间时间戳
	UpdateTimeTimestamp *int64 `json:"update_time_timestamp,omitempty"`

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`
}

func (o ProjectTestCaseFieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectTestCaseFieldVo struct{}"
	}

	return strings.Join([]string{"ProjectTestCaseFieldVo", string(data)}, " ")
}
