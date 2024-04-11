package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteCompareJobRequest Request Object
type DeleteCompareJobRequest struct {

	// 请求语言类型。
	XLanguage *DeleteCompareJobRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务ID。
	CompareJobId string `json:"compare_job_id"`
}

func (o DeleteCompareJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCompareJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteCompareJobRequest", string(data)}, " ")
}

type DeleteCompareJobRequestXLanguage struct {
	value string
}

type DeleteCompareJobRequestXLanguageEnum struct {
	EN_US DeleteCompareJobRequestXLanguage
	ZH_CN DeleteCompareJobRequestXLanguage
}

func GetDeleteCompareJobRequestXLanguageEnum() DeleteCompareJobRequestXLanguageEnum {
	return DeleteCompareJobRequestXLanguageEnum{
		EN_US: DeleteCompareJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteCompareJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DeleteCompareJobRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteCompareJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteCompareJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
