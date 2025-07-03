package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TemplateBlockDeviceMappingOption struct {

	// 虚拟机卷数据源类型：blank, image_id
	SourceId *string `json:"source_id,omitempty"`

	// 卷设备源头类型
	SourceType *TemplateBlockDeviceMappingOptionSourceType `json:"source_type,omitempty"`

	// 是否加密
	Encrypted *bool `json:"encrypted,omitempty"`

	// 秘钥ID
	CmkId *string `json:"cmk_id,omitempty"`

	// 卷类型
	VolumeType *string `json:"volume_type,omitempty"`

	// 卷大小
	VolumeSize *int32 `json:"volume_size,omitempty"`

	Attachment *TemplateBlockDeviceMappingAttachmentOption `json:"attachment,omitempty"`
}

func (o TemplateBlockDeviceMappingOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateBlockDeviceMappingOption struct{}"
	}

	return strings.Join([]string{"TemplateBlockDeviceMappingOption", string(data)}, " ")
}

type TemplateBlockDeviceMappingOptionSourceType struct {
	value string
}

type TemplateBlockDeviceMappingOptionSourceTypeEnum struct {
	BLANK TemplateBlockDeviceMappingOptionSourceType
	IMAGE TemplateBlockDeviceMappingOptionSourceType
}

func GetTemplateBlockDeviceMappingOptionSourceTypeEnum() TemplateBlockDeviceMappingOptionSourceTypeEnum {
	return TemplateBlockDeviceMappingOptionSourceTypeEnum{
		BLANK: TemplateBlockDeviceMappingOptionSourceType{
			value: "blank",
		},
		IMAGE: TemplateBlockDeviceMappingOptionSourceType{
			value: "image",
		},
	}
}

func (c TemplateBlockDeviceMappingOptionSourceType) Value() string {
	return c.value
}

func (c TemplateBlockDeviceMappingOptionSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateBlockDeviceMappingOptionSourceType) UnmarshalJSON(b []byte) error {
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
