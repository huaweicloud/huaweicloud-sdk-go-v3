package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateDataProgressRequest Request Object
type UpdateDataProgressRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *UpdateDataProgressRequestXLanguage `json:"X-Language,omitempty"`

	Body *DataProcessReq `json:"body,omitempty"`
}

func (o UpdateDataProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataProgressRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataProgressRequest", string(data)}, " ")
}

type UpdateDataProgressRequestXLanguage struct {
	value string
}

type UpdateDataProgressRequestXLanguageEnum struct {
	EN_US UpdateDataProgressRequestXLanguage
	ZH_CN UpdateDataProgressRequestXLanguage
}

func GetUpdateDataProgressRequestXLanguageEnum() UpdateDataProgressRequestXLanguageEnum {
	return UpdateDataProgressRequestXLanguageEnum{
		EN_US: UpdateDataProgressRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateDataProgressRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateDataProgressRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateDataProgressRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDataProgressRequestXLanguage) UnmarshalJSON(b []byte) error {
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
