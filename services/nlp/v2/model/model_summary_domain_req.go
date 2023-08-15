package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SummaryDomainReq This is a auto create Body Object
type SummaryDomainReq struct {

	// 生成摘要的长度限制。length_limit > 1，则返回结果为字数不小于该值且最接近该值的摘要。 0 <= length_limit <= 1，则返回结果为长度百分比不小于该值且最接近该值的摘要。
	LengthLimit *float32 `json:"length_limit,omitempty"`

	// 文本标题（目前仅支持UTF-8编码），长度不超过1000字。
	Title *string `json:"title,omitempty"`

	// 支持的文本语言类型，目前支持中文（zh）。
	Lang *SummaryDomainReqLang `json:"lang,omitempty"`

	// 文本正文（目前仅支持UTF-8编码），长度不超过10000字。
	Content string `json:"content"`

	// 支持的领域类型，取值如下（目前只支持通用领域），默认为通用领域： 0：通用领域
	Type *SummaryDomainReqType `json:"type,omitempty"`
}

func (o SummaryDomainReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SummaryDomainReq struct{}"
	}

	return strings.Join([]string{"SummaryDomainReq", string(data)}, " ")
}

type SummaryDomainReqLang struct {
	value string
}

type SummaryDomainReqLangEnum struct {
	ZH SummaryDomainReqLang
}

func GetSummaryDomainReqLangEnum() SummaryDomainReqLangEnum {
	return SummaryDomainReqLangEnum{
		ZH: SummaryDomainReqLang{
			value: "zh",
		},
	}
}

func (c SummaryDomainReqLang) Value() string {
	return c.value
}

func (c SummaryDomainReqLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SummaryDomainReqLang) UnmarshalJSON(b []byte) error {
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

type SummaryDomainReqType struct {
	value int32
}

type SummaryDomainReqTypeEnum struct {
	E_0 SummaryDomainReqType
}

func GetSummaryDomainReqTypeEnum() SummaryDomainReqTypeEnum {
	return SummaryDomainReqTypeEnum{
		E_0: SummaryDomainReqType{
			value: 0,
		},
	}
}

func (c SummaryDomainReqType) Value() int32 {
	return c.value
}

func (c SummaryDomainReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SummaryDomainReqType) UnmarshalJSON(b []byte) error {
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
