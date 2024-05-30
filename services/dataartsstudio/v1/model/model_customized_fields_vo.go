package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CustomizedFieldsVo 自定义项信息。
type CustomizedFieldsVo struct {

	// 编码，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 自定义项中文名称。
	NameCh string `json:"name_ch"`

	// 自定义项英文名称。
	NameEn string `json:"name_en"`

	// 是否必填。
	NotNull bool `json:"not_null"`

	// 可选值。当可选值有多个时，用分号分隔。
	OptionalValues *string `json:"optional_values,omitempty"`

	// 自定义项类型。 枚举值：   - TABLE: 表自定义项   - ATTRIBUTE: 属性自定义项   - SUBJECT: 主题自定义项   - METRIC: 业务指标自定义项
	Type CustomizedFieldsVoType `json:"type"`

	// 系统排序字段，新建、修改时不需要填写。
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 自定义项描述。
	Description *string `json:"description,omitempty"`

	// 创建人，只读。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人，只读。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o CustomizedFieldsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizedFieldsVo struct{}"
	}

	return strings.Join([]string{"CustomizedFieldsVo", string(data)}, " ")
}

type CustomizedFieldsVoType struct {
	value string
}

type CustomizedFieldsVoTypeEnum struct {
	TABLE     CustomizedFieldsVoType
	ATTRIBUTE CustomizedFieldsVoType
	SUBJECT   CustomizedFieldsVoType
	METRIC    CustomizedFieldsVoType
}

func GetCustomizedFieldsVoTypeEnum() CustomizedFieldsVoTypeEnum {
	return CustomizedFieldsVoTypeEnum{
		TABLE: CustomizedFieldsVoType{
			value: "TABLE",
		},
		ATTRIBUTE: CustomizedFieldsVoType{
			value: "ATTRIBUTE",
		},
		SUBJECT: CustomizedFieldsVoType{
			value: "SUBJECT",
		},
		METRIC: CustomizedFieldsVoType{
			value: "METRIC",
		},
	}
}

func (c CustomizedFieldsVoType) Value() string {
	return c.value
}

func (c CustomizedFieldsVoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomizedFieldsVoType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
