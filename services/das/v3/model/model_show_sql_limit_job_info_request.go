package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSqlLimitJobInfoRequest Request Object
type ShowSqlLimitJobInfoRequest struct {
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *ShowSqlLimitJobInfoRequestXLanguage `json:"X-Language,omitempty"`

	// SQL限流任务ID
	JobId string `json:"job_id"`
}

func (o ShowSqlLimitJobInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlLimitJobInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlLimitJobInfoRequest", string(data)}, " ")
}

type ShowSqlLimitJobInfoRequestXLanguage struct {
	value string
}

type ShowSqlLimitJobInfoRequestXLanguageEnum struct {
	ZH_CN ShowSqlLimitJobInfoRequestXLanguage
	EN_US ShowSqlLimitJobInfoRequestXLanguage
}

func GetShowSqlLimitJobInfoRequestXLanguageEnum() ShowSqlLimitJobInfoRequestXLanguageEnum {
	return ShowSqlLimitJobInfoRequestXLanguageEnum{
		ZH_CN: ShowSqlLimitJobInfoRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSqlLimitJobInfoRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSqlLimitJobInfoRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSqlLimitJobInfoRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlLimitJobInfoRequestXLanguage) UnmarshalJSON(b []byte) error {
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
