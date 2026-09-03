package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDownloadUrlRequest Request Object
type CreateDownloadUrlRequest struct {

	// 语言，用于国际化。 - en-us：英文 - zh-cn：中文
	XLanguage *CreateDownloadUrlRequestXLanguage `json:"X-Language,omitempty"`

	// 幂等性标识，UUID格式。 创建类接口携带该请求头，服务端据此实现幂等控制；响应头返回相同值。
	XClientToken *string `json:"X-Client-Token,omitempty"`

	// 技能标识。
	SkillId string `json:"skill_id"`

	// 技能包标识。
	PackageId string `json:"package_id"`

	Body *CreateDownloadUrlReq `json:"body,omitempty"`
}

func (o CreateDownloadUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDownloadUrlRequest struct{}"
	}

	return strings.Join([]string{"CreateDownloadUrlRequest", string(data)}, " ")
}

type CreateDownloadUrlRequestXLanguage struct {
	value string
}

type CreateDownloadUrlRequestXLanguageEnum struct {
	EN_US CreateDownloadUrlRequestXLanguage
	ZH_CN CreateDownloadUrlRequestXLanguage
}

func GetCreateDownloadUrlRequestXLanguageEnum() CreateDownloadUrlRequestXLanguageEnum {
	return CreateDownloadUrlRequestXLanguageEnum{
		EN_US: CreateDownloadUrlRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateDownloadUrlRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateDownloadUrlRequestXLanguage) Value() string {
	return c.value
}

func (c CreateDownloadUrlRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDownloadUrlRequestXLanguage) UnmarshalJSON(b []byte) error {
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
