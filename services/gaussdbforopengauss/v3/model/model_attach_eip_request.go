package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AttachEipRequest Request Object
type AttachEipRequest struct {

	// 语言
	XLanguage *AttachEipRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// 节点ID
	NodeId string `json:"node_id"`

	Body *BindEipRequestBody `json:"body,omitempty"`
}

func (o AttachEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipRequest struct{}"
	}

	return strings.Join([]string{"AttachEipRequest", string(data)}, " ")
}

type AttachEipRequestXLanguage struct {
	value string
}

type AttachEipRequestXLanguageEnum struct {
	ZH_CN AttachEipRequestXLanguage
	EN_US AttachEipRequestXLanguage
}

func GetAttachEipRequestXLanguageEnum() AttachEipRequestXLanguageEnum {
	return AttachEipRequestXLanguageEnum{
		ZH_CN: AttachEipRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: AttachEipRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c AttachEipRequestXLanguage) Value() string {
	return c.value
}

func (c AttachEipRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AttachEipRequestXLanguage) UnmarshalJSON(b []byte) error {
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
