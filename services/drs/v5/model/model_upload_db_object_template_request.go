package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UploadDbObjectTemplateRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *UploadDbObjectTemplateRequestXLanguage `json:"X-Language,omitempty"`

	Body *UploadDbObjectTemplateRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadDbObjectTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadDbObjectTemplateRequest struct{}"
	}

	return strings.Join([]string{"UploadDbObjectTemplateRequest", string(data)}, " ")
}

type UploadDbObjectTemplateRequestXLanguage struct {
	value string
}

type UploadDbObjectTemplateRequestXLanguageEnum struct {
	EN_US UploadDbObjectTemplateRequestXLanguage
	ZH_CN UploadDbObjectTemplateRequestXLanguage
}

func GetUploadDbObjectTemplateRequestXLanguageEnum() UploadDbObjectTemplateRequestXLanguageEnum {
	return UploadDbObjectTemplateRequestXLanguageEnum{
		EN_US: UploadDbObjectTemplateRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UploadDbObjectTemplateRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UploadDbObjectTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c UploadDbObjectTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UploadDbObjectTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
