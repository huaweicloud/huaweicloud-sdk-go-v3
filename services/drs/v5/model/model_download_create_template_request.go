package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DownloadCreateTemplateRequest Request Object
type DownloadCreateTemplateRequest struct {

	// 请求语言类型。
	XLanguage *DownloadCreateTemplateRequestXLanguage `json:"X-Language,omitempty"`

	Body *ExportFilesReq `json:"body,omitempty"`
}

func (o DownloadCreateTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadCreateTemplateRequest struct{}"
	}

	return strings.Join([]string{"DownloadCreateTemplateRequest", string(data)}, " ")
}

type DownloadCreateTemplateRequestXLanguage struct {
	value string
}

type DownloadCreateTemplateRequestXLanguageEnum struct {
	EN_US DownloadCreateTemplateRequestXLanguage
	ZH_CN DownloadCreateTemplateRequestXLanguage
}

func GetDownloadCreateTemplateRequestXLanguageEnum() DownloadCreateTemplateRequestXLanguageEnum {
	return DownloadCreateTemplateRequestXLanguageEnum{
		EN_US: DownloadCreateTemplateRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DownloadCreateTemplateRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DownloadCreateTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c DownloadCreateTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadCreateTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
