package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TransTemplateGroup struct {
	// 模板组名称<br/>

	Name *string `json:"name,omitempty"`
	// 是否设置默认<br/>

	Status *TransTemplateGroupStatus `json:"status,omitempty"`
	// 模板组类型<br/>

	Type *TransTemplateGroupType `json:"type,omitempty"`
	// 是否自动加密

	AutoEncrypt *TransTemplateGroupAutoEncrypt `json:"auto_encrypt,omitempty"`
	// 画质配置信息列表<br/>

	QualityInfoList *[]QualityInfo `json:"quality_info_list,omitempty"`

	Common *Common `json:"common,omitempty"`
	// 绑定的水印模板组ID数组<br/>

	WatermarkTemplateIds *[]string `json:"watermark_template_ids,omitempty"`
	// 模板介绍<br/>

	Description *string `json:"description,omitempty"`
}

func (o TransTemplateGroup) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TransTemplateGroup struct{}"
	}

	return strings.Join([]string{"TransTemplateGroup", string(data)}, " ")
}

type TransTemplateGroupStatus struct {
	value string
}

type TransTemplateGroupStatusEnum struct {
	E_1 TransTemplateGroupStatus
	E_0 TransTemplateGroupStatus
}

func GetTransTemplateGroupStatusEnum() TransTemplateGroupStatusEnum {
	return TransTemplateGroupStatusEnum{
		E_1: TransTemplateGroupStatus{
			value: "1",
		},
		E_0: TransTemplateGroupStatus{
			value: "0",
		},
	}
}

func (c TransTemplateGroupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *TransTemplateGroupStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type TransTemplateGroupType struct {
	value string
}

type TransTemplateGroupTypeEnum struct {
	CUSTOM_TEMPLATE_GROUP TransTemplateGroupType
}

func GetTransTemplateGroupTypeEnum() TransTemplateGroupTypeEnum {
	return TransTemplateGroupTypeEnum{
		CUSTOM_TEMPLATE_GROUP: TransTemplateGroupType{
			value: "custom_template_group",
		},
	}
}

func (c TransTemplateGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *TransTemplateGroupType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type TransTemplateGroupAutoEncrypt struct {
	value int32
}

type TransTemplateGroupAutoEncryptEnum struct {
	E_0 TransTemplateGroupAutoEncrypt
	E_1 TransTemplateGroupAutoEncrypt
}

func GetTransTemplateGroupAutoEncryptEnum() TransTemplateGroupAutoEncryptEnum {
	return TransTemplateGroupAutoEncryptEnum{
		E_0: TransTemplateGroupAutoEncrypt{
			value: 0,
		}, E_1: TransTemplateGroupAutoEncrypt{
			value: 1,
		},
	}
}

func (c TransTemplateGroupAutoEncrypt) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *TransTemplateGroupAutoEncrypt) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
