package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueryTuningParamsRequest Request Object
type QueryTuningParamsRequest struct {

	// 任务ID，如果是主备任务，使用父任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型
	XLanguage *QueryTuningParamsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o QueryTuningParamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTuningParamsRequest struct{}"
	}

	return strings.Join([]string{"QueryTuningParamsRequest", string(data)}, " ")
}

type QueryTuningParamsRequestXLanguage struct {
	value string
}

type QueryTuningParamsRequestXLanguageEnum struct {
	EN_US QueryTuningParamsRequestXLanguage
	ZH_CN QueryTuningParamsRequestXLanguage
}

func GetQueryTuningParamsRequestXLanguageEnum() QueryTuningParamsRequestXLanguageEnum {
	return QueryTuningParamsRequestXLanguageEnum{
		EN_US: QueryTuningParamsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: QueryTuningParamsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c QueryTuningParamsRequestXLanguage) Value() string {
	return c.value
}

func (c QueryTuningParamsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryTuningParamsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
