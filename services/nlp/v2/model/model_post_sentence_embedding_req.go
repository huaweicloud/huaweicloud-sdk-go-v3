package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 命名实体识别post请求体
type PostSentenceEmbeddingReq struct {

	// 文本列表，文本长度为1~512，列表大小为1~1000，文本编码为UTF-8。
	Sentences []string `json:"sentences" xml:"sentences"`

	// 支持的领域类型，目前只支持通用领域，默认为general。
	Domain *PostSentenceEmbeddingReqDomain `json:"domain,omitempty" xml:"domain"`
}

func (o PostSentenceEmbeddingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostSentenceEmbeddingReq struct{}"
	}

	return strings.Join([]string{"PostSentenceEmbeddingReq", string(data)}, " ")
}

type PostSentenceEmbeddingReqDomain struct {
	value string
}

type PostSentenceEmbeddingReqDomainEnum struct {
	GENERAL PostSentenceEmbeddingReqDomain
}

func GetPostSentenceEmbeddingReqDomainEnum() PostSentenceEmbeddingReqDomainEnum {
	return PostSentenceEmbeddingReqDomainEnum{
		GENERAL: PostSentenceEmbeddingReqDomain{
			value: "general",
		},
	}
}

func (c PostSentenceEmbeddingReqDomain) Value() string {
	return c.value
}

func (c PostSentenceEmbeddingReqDomain) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostSentenceEmbeddingReqDomain) UnmarshalJSON(b []byte) error {
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
