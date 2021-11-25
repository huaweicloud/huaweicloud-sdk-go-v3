package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 命名实体识别post请求体
type PostDomainNerRequest struct {
	// 待分析文本，长度为1~64，文本编码为UTF-8。

	Text string `json:"text"`
	// 支持的文本语言类型，目前只支持中文，默认为zh。

	Lang *PostDomainNerRequestLang `json:"lang,omitempty"`
	// 支持的领域类型，目前支持通用（general）领域、商务（business）领域、娱乐（entertainment）领域，默认为general。

	Domain *PostDomainNerRequestDomain `json:"domain,omitempty"`
}

func (o PostDomainNerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostDomainNerRequest struct{}"
	}

	return strings.Join([]string{"PostDomainNerRequest", string(data)}, " ")
}

type PostDomainNerRequestLang struct {
	value string
}

type PostDomainNerRequestLangEnum struct {
	ZH PostDomainNerRequestLang
}

func GetPostDomainNerRequestLangEnum() PostDomainNerRequestLangEnum {
	return PostDomainNerRequestLangEnum{
		ZH: PostDomainNerRequestLang{
			value: "zh",
		},
	}
}

func (c PostDomainNerRequestLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostDomainNerRequestLang) UnmarshalJSON(b []byte) error {
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

type PostDomainNerRequestDomain struct {
	value string
}

type PostDomainNerRequestDomainEnum struct {
	GENERAL       PostDomainNerRequestDomain
	BUSINESS      PostDomainNerRequestDomain
	ENTERTAINMENT PostDomainNerRequestDomain
}

func GetPostDomainNerRequestDomainEnum() PostDomainNerRequestDomainEnum {
	return PostDomainNerRequestDomainEnum{
		GENERAL: PostDomainNerRequestDomain{
			value: "general",
		},
		BUSINESS: PostDomainNerRequestDomain{
			value: "business",
		},
		ENTERTAINMENT: PostDomainNerRequestDomain{
			value: "entertainment",
		},
	}
}

func (c PostDomainNerRequestDomain) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostDomainNerRequestDomain) UnmarshalJSON(b []byte) error {
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
