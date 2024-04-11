package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDataLevelTableCompareJobRequest Request Object
type CreateDataLevelTableCompareJobRequest struct {

	// 请求语言类型。
	XLanguage *CreateDataLevelTableCompareJobRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	Body *CreateDataLevelTableCompareJobReq `json:"body,omitempty"`
}

func (o CreateDataLevelTableCompareJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataLevelTableCompareJobRequest struct{}"
	}

	return strings.Join([]string{"CreateDataLevelTableCompareJobRequest", string(data)}, " ")
}

type CreateDataLevelTableCompareJobRequestXLanguage struct {
	value string
}

type CreateDataLevelTableCompareJobRequestXLanguageEnum struct {
	EN_US CreateDataLevelTableCompareJobRequestXLanguage
	ZH_CN CreateDataLevelTableCompareJobRequestXLanguage
}

func GetCreateDataLevelTableCompareJobRequestXLanguageEnum() CreateDataLevelTableCompareJobRequestXLanguageEnum {
	return CreateDataLevelTableCompareJobRequestXLanguageEnum{
		EN_US: CreateDataLevelTableCompareJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateDataLevelTableCompareJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateDataLevelTableCompareJobRequestXLanguage) Value() string {
	return c.value
}

func (c CreateDataLevelTableCompareJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDataLevelTableCompareJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
