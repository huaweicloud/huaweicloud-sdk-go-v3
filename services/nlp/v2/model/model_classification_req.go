package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ClassificationReq
type ClassificationReq struct {

	// 待分析文本。文本编码要求为utf-8。 限定400个字符以内，文本长度超过400个字符时，只检测前400个字符。
	Content string `json:"content"`

	// 1 广告检测
	Domain *ClassificationReqDomain `json:"domain,omitempty"`
}

func (o ClassificationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClassificationReq struct{}"
	}

	return strings.Join([]string{"ClassificationReq", string(data)}, " ")
}

type ClassificationReqDomain struct {
	value int32
}

type ClassificationReqDomainEnum struct {
	E_1 ClassificationReqDomain
}

func GetClassificationReqDomainEnum() ClassificationReqDomainEnum {
	return ClassificationReqDomainEnum{
		E_1: ClassificationReqDomain{
			value: 1,
		},
	}
}

func (c ClassificationReqDomain) Value() int32 {
	return c.value
}

func (c ClassificationReqDomain) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClassificationReqDomain) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
