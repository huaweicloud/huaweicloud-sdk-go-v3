package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ModifyTransTemplateGroup struct {
	// 模板组名称<br/>

	GroupId *string `json:"group_id,omitempty"`
	// 模板组名称<br/>

	Name *string `json:"name,omitempty"`
	// 是否设置默认<br/>

	Status *ModifyTransTemplateGroupStatus `json:"status,omitempty"`
	// 是否自动加密

	AutoEncrypt *ModifyTransTemplateGroupAutoEncrypt `json:"auto_encrypt,omitempty"`
	// 画质配置信息列表<br/>

	QualityInfoList *[]QualityInfo `json:"quality_info_list,omitempty"`
	// 绑定的水印模板组ID数组<br/>

	WatermarkTemplateIds *[]string `json:"watermark_template_ids,omitempty"`
	// 模板介绍<br/>

	Description *string `json:"description,omitempty"`

	Common *Common `json:"common,omitempty"`
}

func (o ModifyTransTemplateGroup) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyTransTemplateGroup struct{}"
	}

	return strings.Join([]string{"ModifyTransTemplateGroup", string(data)}, " ")
}

type ModifyTransTemplateGroupStatus struct {
	value string
}

type ModifyTransTemplateGroupStatusEnum struct {
	E_1 ModifyTransTemplateGroupStatus
	E_0 ModifyTransTemplateGroupStatus
}

func GetModifyTransTemplateGroupStatusEnum() ModifyTransTemplateGroupStatusEnum {
	return ModifyTransTemplateGroupStatusEnum{
		E_1: ModifyTransTemplateGroupStatus{
			value: "1",
		},
		E_0: ModifyTransTemplateGroupStatus{
			value: "0",
		},
	}
}

func (c ModifyTransTemplateGroupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ModifyTransTemplateGroupStatus) UnmarshalJSON(b []byte) error {
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

type ModifyTransTemplateGroupAutoEncrypt struct {
	value int32
}

type ModifyTransTemplateGroupAutoEncryptEnum struct {
	E_0 ModifyTransTemplateGroupAutoEncrypt
	E_1 ModifyTransTemplateGroupAutoEncrypt
}

func GetModifyTransTemplateGroupAutoEncryptEnum() ModifyTransTemplateGroupAutoEncryptEnum {
	return ModifyTransTemplateGroupAutoEncryptEnum{
		E_0: ModifyTransTemplateGroupAutoEncrypt{
			value: 0,
		}, E_1: ModifyTransTemplateGroupAutoEncrypt{
			value: 1,
		},
	}
}

func (c ModifyTransTemplateGroupAutoEncrypt) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ModifyTransTemplateGroupAutoEncrypt) UnmarshalJSON(b []byte) error {
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
