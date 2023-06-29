package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DownloadDbObjectTemplateRequest Request Object
type DownloadDbObjectTemplateRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *DownloadDbObjectTemplateRequestXLanguage `json:"X-Language,omitempty"`
}

func (o DownloadDbObjectTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDbObjectTemplateRequest struct{}"
	}

	return strings.Join([]string{"DownloadDbObjectTemplateRequest", string(data)}, " ")
}

type DownloadDbObjectTemplateRequestXLanguage struct {
	value string
}

type DownloadDbObjectTemplateRequestXLanguageEnum struct {
	EN_US DownloadDbObjectTemplateRequestXLanguage
	ZH_CN DownloadDbObjectTemplateRequestXLanguage
}

func GetDownloadDbObjectTemplateRequestXLanguageEnum() DownloadDbObjectTemplateRequestXLanguageEnum {
	return DownloadDbObjectTemplateRequestXLanguageEnum{
		EN_US: DownloadDbObjectTemplateRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DownloadDbObjectTemplateRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DownloadDbObjectTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c DownloadDbObjectTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadDbObjectTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
