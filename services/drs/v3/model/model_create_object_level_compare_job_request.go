package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateObjectLevelCompareJobRequest Request Object
type CreateObjectLevelCompareJobRequest struct {

	// 请求语言类型。
	XLanguage *CreateObjectLevelCompareJobRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	Body *CreateObjectCompareJobReq `json:"body,omitempty"`
}

func (o CreateObjectLevelCompareJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObjectLevelCompareJobRequest struct{}"
	}

	return strings.Join([]string{"CreateObjectLevelCompareJobRequest", string(data)}, " ")
}

type CreateObjectLevelCompareJobRequestXLanguage struct {
	value string
}

type CreateObjectLevelCompareJobRequestXLanguageEnum struct {
	EN_US CreateObjectLevelCompareJobRequestXLanguage
	ZH_CN CreateObjectLevelCompareJobRequestXLanguage
}

func GetCreateObjectLevelCompareJobRequestXLanguageEnum() CreateObjectLevelCompareJobRequestXLanguageEnum {
	return CreateObjectLevelCompareJobRequestXLanguageEnum{
		EN_US: CreateObjectLevelCompareJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateObjectLevelCompareJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateObjectLevelCompareJobRequestXLanguage) Value() string {
	return c.value
}

func (c CreateObjectLevelCompareJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateObjectLevelCompareJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
