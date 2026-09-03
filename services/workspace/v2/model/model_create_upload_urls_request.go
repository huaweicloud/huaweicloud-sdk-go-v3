package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateUploadUrlsRequest Request Object
type CreateUploadUrlsRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *CreateUploadUrlsRequestXLanguage `json:"X-Language,omitempty"`

	// 幂等性标识，UUID格式。 创建类接口携带该请求头，服务端据此实现幂等控制；响应头返回相同值。
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateUploadUrlsReq `json:"body,omitempty"`
}

func (o CreateUploadUrlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUploadUrlsRequest struct{}"
	}

	return strings.Join([]string{"CreateUploadUrlsRequest", string(data)}, " ")
}

type CreateUploadUrlsRequestXLanguage struct {
	value string
}

type CreateUploadUrlsRequestXLanguageEnum struct {
	EN_US CreateUploadUrlsRequestXLanguage
	ZH_CN CreateUploadUrlsRequestXLanguage
}

func GetCreateUploadUrlsRequestXLanguageEnum() CreateUploadUrlsRequestXLanguageEnum {
	return CreateUploadUrlsRequestXLanguageEnum{
		EN_US: CreateUploadUrlsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateUploadUrlsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateUploadUrlsRequestXLanguage) Value() string {
	return c.value
}

func (c CreateUploadUrlsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateUploadUrlsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
