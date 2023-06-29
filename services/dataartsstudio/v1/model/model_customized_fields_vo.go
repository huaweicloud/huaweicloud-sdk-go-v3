package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CustomizedFieldsVo 自定义项信息
type CustomizedFieldsVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 中文名称
	NameCh string `json:"name_ch"`

	// 英文名称
	NameEn string `json:"name_en"`

	// 是否必填
	NotNull bool `json:"not_null"`

	// 可选值。分号分隔
	OptionalValues *string `json:"optional_values,omitempty"`

	// 自定义项类型：TABLE, ATTRIBUTE, SUBJECT, METRIC
	Type CustomizedFieldsVoType `json:"type"`

	// 顺序
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
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
