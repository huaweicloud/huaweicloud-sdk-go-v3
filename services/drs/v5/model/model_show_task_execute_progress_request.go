package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTaskExecuteProgressRequest Request Object
type ShowTaskExecuteProgressRequest struct {

	// 请求语言类型。
	XLanguage *ShowTaskExecuteProgressRequestXLanguage `json:"X-Language,omitempty"`

	// 异步任务的唯一标识符。调用导出创建模板接口后从响应中获取该ID，用于后续查询任务执行进度。
	AsyncJobId string `json:"async_job_id"`

	// 查询异步任务类型
	Type *string `json:"type,omitempty"`
}

func (o ShowTaskExecuteProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskExecuteProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskExecuteProgressRequest", string(data)}, " ")
}

type ShowTaskExecuteProgressRequestXLanguage struct {
	value string
}

type ShowTaskExecuteProgressRequestXLanguageEnum struct {
	EN_US ShowTaskExecuteProgressRequestXLanguage
	ZH_CN ShowTaskExecuteProgressRequestXLanguage
	ES_US ShowTaskExecuteProgressRequestXLanguage
	PT_BR ShowTaskExecuteProgressRequestXLanguage
}

func GetShowTaskExecuteProgressRequestXLanguageEnum() ShowTaskExecuteProgressRequestXLanguageEnum {
	return ShowTaskExecuteProgressRequestXLanguageEnum{
		EN_US: ShowTaskExecuteProgressRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowTaskExecuteProgressRequestXLanguage{
			value: "zh-cn",
		},
		ES_US: ShowTaskExecuteProgressRequestXLanguage{
			value: "es-us",
		},
		PT_BR: ShowTaskExecuteProgressRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c ShowTaskExecuteProgressRequestXLanguage) Value() string {
	return c.value
}

func (c ShowTaskExecuteProgressRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTaskExecuteProgressRequestXLanguage) UnmarshalJSON(b []byte) error {
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
