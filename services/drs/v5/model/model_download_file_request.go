package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DownloadFileRequest Request Object
type DownloadFileRequest struct {

	// 请求语言类型。
	XLanguage *DownloadFileRequestXLanguage `json:"X-Language,omitempty"`

	Body *ExportFilesReq `json:"body,omitempty"`
}

func (o DownloadFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadFileRequest struct{}"
	}

	return strings.Join([]string{"DownloadFileRequest", string(data)}, " ")
}

type DownloadFileRequestXLanguage struct {
	value string
}

type DownloadFileRequestXLanguageEnum struct {
	EN_US DownloadFileRequestXLanguage
	ZH_CN DownloadFileRequestXLanguage
	ES_US DownloadFileRequestXLanguage
	PT_BR DownloadFileRequestXLanguage
}

func GetDownloadFileRequestXLanguageEnum() DownloadFileRequestXLanguageEnum {
	return DownloadFileRequestXLanguageEnum{
		EN_US: DownloadFileRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DownloadFileRequestXLanguage{
			value: "zh-cn",
		},
		ES_US: DownloadFileRequestXLanguage{
			value: "es-us",
		},
		PT_BR: DownloadFileRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c DownloadFileRequestXLanguage) Value() string {
	return c.value
}

func (c DownloadFileRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadFileRequestXLanguage) UnmarshalJSON(b []byte) error {
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
