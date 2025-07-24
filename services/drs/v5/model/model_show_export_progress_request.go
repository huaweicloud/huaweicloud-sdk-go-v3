package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowExportProgressRequest Request Object
type ShowExportProgressRequest struct {

	// 请求语言类型。
	XLanguage *ShowExportProgressRequestXLanguage `json:"X-Language,omitempty"`

	// 导出创建模板接口返回的异步ID。
	AsyncJobId string `json:"async_job_id"`
}

func (o ShowExportProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExportProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowExportProgressRequest", string(data)}, " ")
}

type ShowExportProgressRequestXLanguage struct {
	value string
}

type ShowExportProgressRequestXLanguageEnum struct {
	EN_US ShowExportProgressRequestXLanguage
	ZH_CN ShowExportProgressRequestXLanguage
}

func GetShowExportProgressRequestXLanguageEnum() ShowExportProgressRequestXLanguageEnum {
	return ShowExportProgressRequestXLanguageEnum{
		EN_US: ShowExportProgressRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowExportProgressRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowExportProgressRequestXLanguage) Value() string {
	return c.value
}

func (c ShowExportProgressRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowExportProgressRequestXLanguage) UnmarshalJSON(b []byte) error {
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
