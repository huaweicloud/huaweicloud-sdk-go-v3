package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TemplateGroup struct {
	// 模板组id<br/>

	GroupId *string `json:"group_id,omitempty"`
	// 模板组名称<br/>

	Name *string `json:"name,omitempty"`
	// 是否默认<br/>

	Status *string `json:"status,omitempty"`
	// 模板组类型<br/>

	Type *string `json:"type,omitempty"`
	// 是否自动加密

	AutoEncrypt *TemplateGroupAutoEncrypt `json:"auto_encrypt,omitempty"`
	// 画质配置信息列表<br/>

	QualityInfoList *[]QualityInfo `json:"quality_info_list,omitempty"`
	// 绑定的水印模板组ID数组<br/>

	WatermarkTemplateIds *[]string `json:"watermark_template_ids,omitempty"`
	// 模板介绍<br/>

	Description *string `json:"description,omitempty"`

	Common *Common `json:"common,omitempty"`
}

func (o TemplateGroup) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TemplateGroup struct{}"
	}

	return strings.Join([]string{"TemplateGroup", string(data)}, " ")
}

type TemplateGroupAutoEncrypt struct {
	value int32
}

type TemplateGroupAutoEncryptEnum struct {
	E_0 TemplateGroupAutoEncrypt
	E_1 TemplateGroupAutoEncrypt
}

func GetTemplateGroupAutoEncryptEnum() TemplateGroupAutoEncryptEnum {
	return TemplateGroupAutoEncryptEnum{
		E_0: TemplateGroupAutoEncrypt{
			value: 0,
		}, E_1: TemplateGroupAutoEncrypt{
			value: 1,
		},
	}
}

func (c TemplateGroupAutoEncrypt) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *TemplateGroupAutoEncrypt) UnmarshalJSON(b []byte) error {
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
