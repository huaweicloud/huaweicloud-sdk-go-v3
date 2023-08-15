package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PostEntityLinkingRequest 实体链接post请求体
type PostEntityLinkingRequest struct {

	// 待分析文本，长度为1~50，文本编码为UTF-8。
	Text string `json:"text"`

	// 支持的文本语言类型，目前只支持中文，默认为zh。
	Lang *PostEntityLinkingRequestLang `json:"lang,omitempty"`
}

func (o PostEntityLinkingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostEntityLinkingRequest struct{}"
	}

	return strings.Join([]string{"PostEntityLinkingRequest", string(data)}, " ")
}

type PostEntityLinkingRequestLang struct {
	value string
}

type PostEntityLinkingRequestLangEnum struct {
	ZH PostEntityLinkingRequestLang
}

func GetPostEntityLinkingRequestLangEnum() PostEntityLinkingRequestLangEnum {
	return PostEntityLinkingRequestLangEnum{
		ZH: PostEntityLinkingRequestLang{
			value: "zh",
		},
	}
}

func (c PostEntityLinkingRequestLang) Value() string {
	return c.value
}

func (c PostEntityLinkingRequestLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostEntityLinkingRequestLang) UnmarshalJSON(b []byte) error {
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
