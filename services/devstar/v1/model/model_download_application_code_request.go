package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DownloadApplicationCodeRequest Request Object
type DownloadApplicationCodeRequest struct {

	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文
	XLanguage *DownloadApplicationCodeRequestXLanguage `json:"X-Language,omitempty"`

	// 任务id。
	JobId string `json:"job_id"`
}

func (o DownloadApplicationCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadApplicationCodeRequest struct{}"
	}

	return strings.Join([]string{"DownloadApplicationCodeRequest", string(data)}, " ")
}

type DownloadApplicationCodeRequestXLanguage struct {
	value string
}

type DownloadApplicationCodeRequestXLanguageEnum struct {
	ZH_CN DownloadApplicationCodeRequestXLanguage
	EN_US DownloadApplicationCodeRequestXLanguage
}

func GetDownloadApplicationCodeRequestXLanguageEnum() DownloadApplicationCodeRequestXLanguageEnum {
	return DownloadApplicationCodeRequestXLanguageEnum{
		ZH_CN: DownloadApplicationCodeRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DownloadApplicationCodeRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DownloadApplicationCodeRequestXLanguage) Value() string {
	return c.value
}

func (c DownloadApplicationCodeRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadApplicationCodeRequestXLanguage) UnmarshalJSON(b []byte) error {
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
